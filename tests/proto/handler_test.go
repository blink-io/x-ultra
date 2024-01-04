package metadata

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"testing"

	grpcg "github.com/blink-io/x/kratos/v2/transport/grpc/g"
	khttp "github.com/blink-io/x/kratos/v2/transport/http"
	httpg "github.com/blink-io/x/kratos/v2/transport/http/g"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	kgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/stretchr/testify/require"
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
	h := grpcg.NewHandler[MetadataXServer](s, RegisterMetadataXServer)
	require.NotNil(t, h)

	ln, err := net.Listen("tcp", ":9997")
	require.NoError(t, err)

	gsrv := kgrpc.NewServer(
		kgrpc.Listener(ln),
		kgrpc.Options(
			grpc.Creds(insecure.NewCredentials()),
		),
	)

	h.HandleGRPC(gsrv)

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
	h := httpg.NewHandler[MetadataXHTTPServer](s, RegisterMetadataXHTTPServer)
	require.NotNil(t, h)

	hsrv := khttp.NewServer(
		khttp.Address(":9996"),
	)

	h.HandleHTTP(hsrv)

	require.NoError(t, hsrv.Start(context.Background()))

	fmt.Println("done")
}

func TestHandler_HTTP_Server_2(t *testing.T) {
	s := new(mmm)
	h := httpg.NewHandler[*mmm](s, RegisterMMM)
	require.NotNil(t, h)

	hsrv := khttp.NewServer(
		khttp.Address(":9996"),
	)

	h.HandleHTTP(hsrv)

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

type hhdlr = httpg.Handler[MetadataXHTTPServer]
type ghdlr = grpcg.Handler[MetadataXServer]

type compose struct {
	hhdlr
	ghdlr
}

func (c *compose) HandleHTTP(r khttp.ServerRouter) {
	c.hhdlr.HandleHTTP(r)
}

func (c *compose) HandleGRPC(r grpc.ServiceRegistrar) {
	c.ghdlr.HandleGRPC(r)
}

type mmm struct {
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

func (h *mmm) GetMyTime() khttp.HandlerFunc {
	f := httpg.GET[Req, Res]("get/do-my-time", handleMyTime)
	return f
}

func (h *mmm) PostMyTime() khttp.HandlerFunc {
	f := httpg.POST[Req, Res]("post/do-my-time", handleMyTime)
	return f
}

func RegisterMMM(r khttp.ServerRouter, h *mmm) {
	sr := r.Route("/mmm")
	sr.POST("/do-my-time", h.PostMyTime())
	sr.POST("/do-my-time", h.GetMyTime())
	sr.GET("/do-my-time/v2", httpg.Func[Req, Res](handleMyTime).Do(http.MethodGet, "get/do-my-time/v2"))
}

func TestInit_1(t *testing.T) {
	var res *Res
	res.Init()
	fmt.Println("")
}

func TestHandler_Compose_1(t *testing.T) {
	s := new(service)
	hh := httpg.NewHandler[MetadataXHTTPServer](s, RegisterMetadataXHTTPServer)
	gh := grpcg.NewHandler[MetadataXServer](s, RegisterMetadataXServer)

	co := &compose{
		hhdlr: hh,
		ghdlr: gh,
	}

	hsrv := khttp.NewServer()
	gsrv := kgrpc.NewServer()

	co.HandleHTTP(hsrv)
	co.HandleGRPC(gsrv)

	fmt.Println("done")
}
