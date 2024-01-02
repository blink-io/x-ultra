package http3

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/blink-io/x/internal/testdata"
	"github.com/blink-io/x/kratos/v2/internal/host"
	"github.com/stretchr/testify/require"

	kerrors "github.com/go-kratos/kratos/v2/errors"
	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/http3"
)

var (
	clientTlsConf = testdata.GetClientTLSConfig()

	serverTlsConf = testdata.GetServerTLSConfig()

	h = func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(testData{Path: r.RequestURI})
	}

	HTTP3Client = &http.Client{
		Timeout:   5 * time.Second,
		Transport: RoundTripper(clientTlsConf),
	}
)

func init() {
	http.DefaultClient = HTTP3Client
}

type testKey struct{}

type testData struct {
	Path string `json:"path"`
}

func TLSConfigClientOption() ClientOption {
	return WithTLSConfig(clientTlsConf)
}

func TLSConfigServerOption() ServerOption {
	return TLSConfig(serverTlsConf)
}

func CreateListener() http3.QUICEarlyListener {
	ln, err := quic.ListenAddrEarly(":0", http3.ConfigureTLSConfig(serverTlsConf), nil)
	if err != nil {
		panic(err)
	}
	return ln
}

// handleFuncWrapper is a wrapper for http.HandlerFunc to implement http.Handler
type handleFuncWrapper struct {
	fn http.HandlerFunc
}

func (x *handleFuncWrapper) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	x.fn.ServeHTTP(writer, request)
}

func newHandleFuncWrapper(fn http.HandlerFunc) http.Handler {
	return &handleFuncWrapper{fn: fn}
}

func TestServeHTTP3(t *testing.T) {
	ln := CreateListener()

	adp := NewAdapter(DefaultOptions,
		Listener(ln),
		QConfig(new(quic.Config)),
	)
	mux := NewServer(
		Adapter(adp),
	)
	mux.HandleFunc("/index", h)
	mux.Route("/errors").GET("/cause", func(ctx Context) error {
		return kerrors.BadRequest("xxx", "zzz").
			WithMetadata(map[string]string{"foo": "bar"}).
			WithCause(errors.New("error cause"))
	})
	if err := mux.WalkRoute(func(r RouteInfo) error {
		t.Logf("WalkRoute: %+v", r)
		return nil
	}); err != nil {
		t.Fatal(err)
	}
	if e, err := mux.Endpoint(); err != nil || e == nil || strings.HasSuffix(e.Host, ":0") {
		t.Fatal(e, err)
	}
	srv := http3.Server{Handler: mux}
	go func() {
		if err := srv.ServeListener(ln); err != nil {
			if kerrors.Is(err, http.ErrServerClosed) {
				return
			}
			panic(err)
		}
	}()
	time.Sleep(time.Second)
	if err := srv.Close(); err != nil {
		t.Log(err)
	}
}

func TestServerHTTP3(t *testing.T) {
	ctx := context.Background()
	srv := NewServer(TLSConfigServerOption())
	srv.Handle("/index", newHandleFuncWrapper(h))
	srv.HandleFunc("/index/{id:[0-9]+}", h)
	srv.HandlePrefix("/test/prefix", newHandleFuncWrapper(h))
	srv.HandleHeader("content-type", "application/grpc-web+json", func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(testData{Path: r.RequestURI})
	})
	srv.Route("/errors").GET("/cause", func(ctx Context) error {
		return kerrors.BadRequest("xxx", "zzz").
			WithMetadata(map[string]string{"foo": "bar"}).
			WithCause(errors.New("error cause"))
	})

	if e, err := srv.Endpoint(); err != nil || e == nil || strings.HasSuffix(e.Host, ":0") {
		t.Fatal(e, err)
	}

	go func() {
		if err := srv.Start(ctx); err != nil {
			panic(err)
		}
	}()
	time.Sleep(time.Second)
	testHeaderHTTP3(t, srv)
	testClientHTTP3(t, srv)
	testAcceptHTTP3(t, srv)
	time.Sleep(time.Second)

	if srv.Stop(ctx) != nil {
		t.Errorf("expected nil got %v", srv.Stop(ctx))
	}
}

func testAcceptHTTP3(t *testing.T, srv Server) {
	tests := []struct {
		method      string
		path        string
		contentType string
	}{
		{http.MethodGet, "/errors/cause", "application/json"},
		{http.MethodGet, "/errors/cause", "application/proto"},
	}
	e, err := srv.Endpoint()
	if err != nil {
		t.Errorf("expected nil got %v", err)
	}
	client, err := NewClient(
		context.Background(),
		WithEndpoint(e.Host),
		WithTransport(RoundTripper(clientTlsConf)),
		WithTLSConfig(clientTlsConf))
	if err != nil {
		t.Errorf("expected nil got %v", err)
	}
	for _, test := range tests {
		req, err := http.NewRequest(test.method, e.String()+test.path, nil)
		if err != nil {
			t.Errorf("expected nil got %v", err)
		}
		req.Header.Set("Content-Type", test.contentType)
		resp, err := client.Do(req)
		if kerrors.Code(err) != 400 {
			t.Errorf("expected 400 got %v", err)
		}
		if err == nil {
			resp.Body.Close()
		}
	}
}

func testHeaderHTTP3(t *testing.T, srv Server) {
	e, err := srv.Endpoint()
	if err != nil {
		t.Errorf("expected nil got %v", err)
	}
	client, err := NewClient(
		context.Background(),
		WithEndpoint(e.Host),
		WithTransport(RoundTripper(clientTlsConf)),
		WithTLSConfig(clientTlsConf),
	)
	if err != nil {
		t.Errorf("expected nil got %v", err)
	}

	reqURL := fmt.Sprintf(e.String() + "/index")
	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		t.Errorf("expected nil got %v", err)
	}
	req.Header.Set("content-type", "application/grpc-web+json")
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("expected nil got %v", err)
	}
	resp.Body.Close()
}

func testClientHTTP3(t *testing.T, srv Server) {
	tests := []struct {
		method string
		path   string
		code   int
	}{
		{http.MethodGet, "/index", http.StatusOK},
		{http.MethodPut, "/index", http.StatusOK},
		{http.MethodPost, "/index", http.StatusOK},
		{http.MethodPatch, "/index", http.StatusOK},
		{http.MethodDelete, "/index", http.StatusOK},

		{http.MethodGet, "/index/1", http.StatusOK},
		{http.MethodPut, "/index/1", http.StatusOK},
		{http.MethodPost, "/index/1", http.StatusOK},
		{http.MethodPatch, "/index/1", http.StatusOK},
		{http.MethodDelete, "/index/1", http.StatusOK},

		{http.MethodGet, "/index/notfound", http.StatusNotFound},
		{http.MethodGet, "/errors/cause", http.StatusBadRequest},
		{http.MethodGet, "/test/prefix/123111", http.StatusOK},
	}
	e, err := srv.Endpoint()
	if err != nil {
		t.Fatal(err)
	}
	client, err := NewClient(
		context.Background(),
		WithEndpoint(e.Host),
		WithTransport(RoundTripperConf(clientTlsConf, nil)),
		WithTLSConfig(clientTlsConf),
	)
	if err != nil {
		t.Fatal(err)
	}

	defer client.Close()
	for _, test := range tests {
		var res testData
		reqURL := fmt.Sprintf(e.String() + test.path)
		req, err := http.NewRequest(test.method, reqURL, nil)
		if err != nil {
			t.Fatal(err)
		}
		resp, err := client.Do(req)
		if kerrors.Code(err) != test.code {
			t.Fatalf("want %v, but got %v", test, err)
		}
		if err != nil {
			continue
		}
		if resp.StatusCode != 200 {
			_ = resp.Body.Close()
			t.Fatalf("http status got %d", resp.StatusCode)
		}
		content, err := io.ReadAll(resp.Body)
		_ = resp.Body.Close()
		if err != nil {
			t.Fatalf("read resp error %v", err)
		}
		err = json.Unmarshal(content, &res)
		if err != nil {
			t.Fatalf("unmarshal resp error %v", err)
		}
		if res.Path != test.path {
			t.Errorf("expected %s got %s", test.path, res.Path)
		}
	}
	for _, test := range tests {
		var res testData
		err := client.Invoke(context.Background(), test.method, test.path, nil, &res)
		if kerrors.Code(err) != test.code {
			t.Fatalf("want %v, but got %v", test, err)
		}
		if err != nil {
			continue
		}
		if res.Path != test.path {
			t.Errorf("expected %s got %s", test.path, res.Path)
		}
	}
}

func BenchmarkServerHTTP3(b *testing.B) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		data := &testData{Path: r.RequestURI}
		_ = json.NewEncoder(w).Encode(data)
		if r.Context().Value(testKey{}) != "test" {
			w.WriteHeader(500)
		}
	}
	ctx := context.Background()
	ctx = context.WithValue(ctx, testKey{}, "test")
	srv := NewServer(TLSConfigServerOption())
	srv.HandleFunc("/index", fn)
	go func() {
		if err := srv.Start(ctx); err != nil {
			panic(err)
		}
	}()
	time.Sleep(time.Second)
	ln := srv.Listener()
	port, ok := host.Port(ln)
	if !ok {
		b.Errorf("expected port got %v", ln)
	}
	client, err := NewClient(
		context.Background(),
		WithEndpoint(fmt.Sprintf("127.0.0.1:%d", port)),
		WithTransport(RoundTripperConf(clientTlsConf, nil)),
		WithTLSConfig(clientTlsConf))
	if err != nil {
		b.Errorf("expected nil got %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var res testData
		err := client.Invoke(context.Background(), http.MethodPost, "/index", nil, &res)
		if err != nil {
			b.Errorf("expected nil got %v", err)
		}
	}
	_ = srv.Stop(ctx)
}

func TestStartServer(t *testing.T) {
	srv := NewServer(TLSConfigServerOption(), Address(":9999"))

	srv.Handle("/hello", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "plain/text")
		w.Write([]byte("Are you OK?"))
	}))
	srv.Start(context.Background())
}

func TestStartClient(t *testing.T) {
	cc, err := NewClient(context.Background(),
		WithTransport(RoundTripperConf(clientTlsConf, nil)),
		WithEndpoint("https://localhost:9999"),
	)
	require.NoError(t, err)

	req, err := http.NewRequest(http.MethodGet, "https://localhost:9999/hello", nil)
	require.NoError(t, err)

	res, err := cc.Do(req)
	require.NoError(t, err)

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	require.NoError(t, err)

	fmt.Println("res: ", string(data))
}
