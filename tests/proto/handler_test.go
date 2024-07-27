package metadata

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"testing"

	kgrpc "github.com/blink-io/x/kratos/v2/transport/grpc"
	khttp "github.com/blink-io/x/kratos/v2/transport/http"
	khttpdo "github.com/blink-io/x/kratos/v2/transport/http/do"

	"connectrpc.com/grpchealth"
	"connectrpc.com/grpcreflect"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	profilingsvc "google.golang.org/grpc/profiling/service"
)

type service struct {
	*UnimplementedMetadataXServer
}

func (s *service) ListServices(ctx context.Context, req *XListServicesRequest) (*XListServicesReply, error) {
	reply := &XListServicesReply{
		Services: []string{
			"s1",
			"s2",
		},
		Methods: []string{
			"m1",
			"m2",
		},
	}
	return reply, nil
}

func (s *service) GetServiceDesc(ctx context.Context, req *XGetServiceDescRequest) (*XGetServiceDescReply, error) {
	reply := &XGetServiceDescReply{
		FileDescSet: nil,
	}
	return reply, nil
}

var _ MetadataXServer = (*service)(nil)
var _ MetadataXHTTPServer = (*service)(nil)

func TestHandler_GRPC_Server_1(t *testing.T) {
	s := new(service)
	CtxRegisterMetadataXServer := func(ctx context.Context, s grpc.ServiceRegistrar, srv MetadataXServer) {
		RegisterMetadataXServer(s, srv)
	}
	mh := kgrpc.NewCtxRegistrar[MetadataXServer](s, CtxRegisterMetadataXServer)
	hh := kgrpc.NewRegistrar[grpc_health_v1.HealthServer](health.NewServer(), grpc_health_v1.RegisterHealthServer)

	ln, err := net.Listen("tcp", ":9997")
	require.NoError(t, err)

	gsrv := kgrpc.NewServer(
		kgrpc.Listener(ln),
		kgrpc.Options(
			grpc.Creds(insecure.NewCredentials()),
		),
	)

	mh.RegisterToGRPC(context.Background(), gsrv)
	hh.RegisterToGRPC(context.Background(), gsrv)

	err2 := profilingsvc.Init(&profilingsvc.ProfilingConfig{
		Enabled: true,
		Server:  gsrv.Raw(),
	})
	require.NoError(t, err2)

	require.NoError(t, gsrv.Start(context.Background()))

	fmt.Println("done")
}

func TestHandler_GRPC_Client_1(t *testing.T) {
	cc, err := kgrpc.DialInsecure(context.Background(),
		kgrpc.WithEndpoint("localhost:9997"),
	)
	require.NoError(t, err)

	cli := NewMetadataXClient(cc)

	reply, err := cli.ListServices(context.Background(), &XListServicesRequest{})
	require.NoError(t, err)

	fmt.Println("Reply: ", reply)
}

func TestHandler_HTTP_Server_1(t *testing.T) {
	s := new(service)
	h := khttp.NewRegistrar[MetadataXHTTPServer](s, RegisterMetadataXHTTPServer)
	require.NotNil(t, h)

	hsrv := khttp.NewServer(
		khttp.Address(":9996"),
	)

	h.RegisterToHTTP(context.Background(), hsrv)

	require.NoError(t, hsrv.Start(context.Background()))

	fmt.Println("done")
}

func TestHandler_HTTP_Server_2(t *testing.T) {
	s := new(MyTimeSvc)
	//h := khttpx.NewRegistrar[*MyTimeSvc](s, RegisterMyTimeSvc)
	//require.NotNil(t, h)

	hsrv := khttp.NewServer(
		khttp.Address(":9996"),
	)

	s.HTTPRegistrar().RegisterToHTTP(context.Background(), hsrv)

	require.NoError(t, hsrv.Start(context.Background()))

	fmt.Println("done")
}

func TestHandler_HTTP_Client_1(t *testing.T) {
	cc, err := khttp.NewClient(context.Background(),
		khttp.WithEndpoint("localhost:9996"),
	)
	require.NoError(t, err)

	cli := NewMetadataXHTTPClient(cc)

	reply, err := cli.ListServices(context.Background(), &XListServicesRequest{})
	require.NoError(t, err)

	fmt.Println("Reply: ", reply)
}

type hhdlr = khttp.Registrar
type ghdlr = kgrpc.Registrar

var _ khttp.Registrar = (*compose)(nil)
var _ kgrpc.Registrar = (*compose)(nil)

type compose struct {
	hhdlr
	ghdlr
}

func (c *compose) HandleHTTP(ctx context.Context, r khttp.ServerRouter) {
	c.hhdlr.RegisterToHTTP(ctx, r)
}

func (c *compose) HandleGRPC(ctx context.Context, r kgrpc.ServiceRegistrar) {
	c.ghdlr.RegisterToGRPC(ctx, r)
}

var _ khttp.WithRegistrar = (*MyTimeSvc)(nil)

type MyTimeSvc struct {
}

func (h *MyTimeSvc) HTTPRegistrar() khttp.Registrar {
	return khttp.NewRegistrar(h, RegisterMyTimeSvc)
}

type Req struct {
	Msg string `json:"msg"`
}

type Res struct {
	Payload string `json:"payload"`
}

func (r Res) Init() {
	fmt.Println("HHHH")
}

func handleMyTime(ctx context.Context, r *Req) (*Res, error) {
	msg := r.Msg
	fmt.Println("msg: ", msg)

	res := &Res{
		Payload: "哈哈哈" + msg,
	}
	return res, nil
}

func (h *MyTimeSvc) GetMyTime() khttp.HandlerFunc {
	f := khttpdo.GET[Req, Res]("get/do-my-time", handleMyTime)
	return f
}

func (h *MyTimeSvc) PostMyTime() khttp.HandlerFunc {
	f := khttpdo.POST[Req, Res]("post/do-my-time", handleMyTime)
	return f
}

func RegisterMyTimeSvc(r khttp.ServerRouter, h *MyTimeSvc) {
	sr := r.Route("/MyTimeSvc")
	sr.POST("/do-my-time", h.PostMyTime())
	sr.GET("/do-my-time", h.GetMyTime())
	sr.GET("/do-my-time/v2", khttpdo.Func[Req, Res](handleMyTime).Do(http.MethodGet, "get:do-my-time/v2"))

	checker := grpchealth.NewStaticChecker(
		"acme.user.v1.UserService",
		"acme.group.v1.GroupService",
		// protoc-gen-connect-go generates package-level constants
		// for these fully-qualified protobuf service names, so you'd more likely
	)
	r.HandlePrefix(grpchealth.NewHandler(checker))

	reflector := grpcreflect.NewStaticReflector(
		"acme.user.v1.UserService",
		"acme.group.v1.GroupService",
		// protoc-gen-connect-go generates package-level constants
		// for these fully-qualified protobuf service names, so you'd more likely
		// reference userv1.UserServiceName and groupv1.GroupServiceName.
	)
	r.HandlePrefix(grpcreflect.NewHandlerV1(reflector))
}

func TestInit_1(t *testing.T) {
	var res *Res
	res.Init()
	fmt.Println("")
}

func TestRegistrar_Compose_1(t *testing.T) {
	s := new(service)
	CtxRegisterMetadataXHTTPServer := func(ctx context.Context, s khttp.ServerRouter, srv MetadataXHTTPServer) {
		RegisterMetadataXHTTPServer(s, srv)
	}
	hh := khttp.NewCtxRegistrar[MetadataXHTTPServer](s, CtxRegisterMetadataXHTTPServer)
	gh := kgrpc.NewRegistrar[MetadataXServer](s, RegisterMetadataXServer)

	co := &compose{
		hhdlr: hh,
		ghdlr: gh,
	}

	ctx := context.Background()
	hsrv := khttp.NewServer()
	gsrv := kgrpc.NewServer()

	co.HandleHTTP(ctx, hsrv)
	co.HandleGRPC(ctx, gsrv)

	fmt.Println("done")
}
