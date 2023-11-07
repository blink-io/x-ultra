package grpc_test

import (
	"context"
	"fmt"
	"testing"

	sessgrpc "github.com/blink-io/x/session/grpc"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func TestGRPC_Client_1(t *testing.T) {
	ctx := context.Background()
	ops := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	c, err := grpc.Dial(":9999", ops...)
	require.NoError(t, err)

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
