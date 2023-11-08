package grpc_test

import (
	"context"
	"fmt"
	"log"
	"testing"

	sessgrpc "github.com/blink-io/x/session/grpc"
	"github.com/blink-io/x/testdata"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

func createGRPCClient() *grpc.ClientConn {
	creds := credentials.NewTLS(testdata.CreateClientTLSConfig())
	ops := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}
	c, err := grpc.Dial(":9999", ops...)
	if err != nil {
		log.Fatal(err)
	}
	return c
}

func TestGRPC_SessMgr_Client_1(t *testing.T) {
	ctx := context.Background()
	c := createGRPCClient()
	cc := NewCommonClient(c)

	req := &HealthRequest{
		From: "我是来自GRPPC Client的Session值",
	}
	var header, trailer metadata.MD
	res, err := cc.Health(ctx, req, grpc.Header(&header), grpc.Trailer(&trailer))
	require.NoError(t, err)

	fmt.Println("Health res:  ", res)
	fmt.Println("header:  ", header)
	fmt.Println("trailer:  ", trailer)

	getFirst := func(ss []string) string {
		if len(ss) > 0 {
			return ss[0]
		}
		return ""
	}
	token := getFirst(header.Get(sessgrpc.DefaultHeader))
	fmt.Println("token:  ", token)

	mctx := metadata.AppendToOutgoingContext(ctx, sessgrpc.DefaultHeader, token)
	vres, verr := cc.Version(mctx, &VersionRequest{
		From: "From_Mama",
	})
	require.NoError(t, verr)
	require.NotNil(t, vres)

	fmt.Println("Version res:  ", vres)
}

func TestMD_1(t *testing.T) {
	md := metadata.Pairs()
	md.Set("a", "111")
	require.NotNil(t, md)
}

func TestGRPC_SessMgr_Client_2(t *testing.T) {
	ctx := context.Background()
	c := createGRPCClient()
	cc := NewCommonClient(c)

	req := &TestingRequest{
		Action: "我是来自GRPPC Client的Testing",
	}
	var header, trailer metadata.MD
	res, err := cc.Testing(ctx, req, grpc.Header(&header), grpc.Trailer(&trailer))
	require.NoError(t, err)

	fmt.Println("Health res:  ", res)
	fmt.Println("header:  ", header)
	fmt.Println("trailer:  ", trailer)
}
