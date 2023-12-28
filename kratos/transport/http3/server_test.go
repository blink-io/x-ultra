package http3

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"testing"

	khttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/kratos/v2/transport/http/binding"

	"github.com/stretchr/testify/assert"

	api "github.com/blink-io/x/internal/testing/api/protobuf"
)

func HygrothermographHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("HygrothermographHandler [%s] [%s] [%s]\n", r.Proto, r.Method, r.RequestURI)

	if r.Method == "POST" {
		var in api.Hygrothermograph
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
			fmt.Printf("decode error: %s\n", err.Error())
		}
		fmt.Printf("Humidity: %s Temperature: %s \n", in.Humidity, in.Temperature)
	}

	var out api.Hygrothermograph
	out.Humidity = strconv.FormatInt(int64(rand.Intn(100)), 10)
	out.Temperature = strconv.FormatInt(int64(rand.Intn(100)), 10)
	_ = json.NewEncoder(w).Encode(&out)
}

func TestServer(t *testing.T) {
	ctx := context.Background()

	srv := NewServer(
		WithAddress(":8800"),
	)

	srv.HandleFunc("/hygrothermograph", HygrothermographHandler)

	if err := srv.Start(ctx); err != nil {
		panic(err)
	}

	defer func() {
		if err := srv.Stop(ctx); err != nil {
			t.Errorf("expected nil got %v", err)
		}
	}()
}

func GetHygrothermograph(ctx context.Context, cli *khttp.Client, in *api.Hygrothermograph, opts ...khttp.CallOption) (*api.Hygrothermograph, error) {
	var out api.Hygrothermograph

	pattern := "/hygrothermograph"
	path := binding.EncodeURL(pattern, in, true)

	opts = append(opts, khttp.Operation("/GetHygrothermograph"))
	opts = append(opts, khttp.PathTemplate(pattern))

	err := cli.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

func CreateHygrothermograph(ctx context.Context, cli *khttp.Client, in *api.Hygrothermograph, opts ...khttp.CallOption) (*api.Hygrothermograph, error) {
	var out api.Hygrothermograph

	pattern := "/hygrothermograph"
	path := binding.EncodeURL(pattern, in, false)

	opts = append(opts, khttp.Operation("/CreateHygrothermograph"))
	opts = append(opts, khttp.PathTemplate(pattern))

	err := cli.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

func TestClient(t *testing.T) {
	ctx := context.Background()

	pool, err := x509.SystemCertPool()
	assert.Nil(t, err)

	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		RootCAs:            pool,
	}
	cli, err := khttp.NewClient(ctx,
		khttp.WithEndpoint("127.0.0.1:8800"),
		khttp.WithTLSConfig(tlsConf),
		khttp.WithTransport(RoundTripper(tlsConf)),
	)
	assert.Nil(t, err)
	assert.NotNil(t, cli)

	var req api.Hygrothermograph
	req.Humidity = strconv.FormatInt(int64(rand.Intn(100)), 10)
	req.Temperature = strconv.FormatInt(int64(rand.Intn(100)), 10)

	resp, err := GetHygrothermograph(ctx, cli, &req, khttp.EmptyCallOption{})
	assert.Nil(t, err)
	t.Log(resp)

	resp, err = CreateHygrothermograph(ctx, cli, &req, khttp.EmptyCallOption{})
	assert.Nil(t, err)
	t.Log(resp)
}
