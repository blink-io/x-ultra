package requestid

import (
	"context"

	"github.com/blink-io/x/grpc/mdutil"
	"github.com/blink-io/x/grpc/util"
	"github.com/blink-io/x/requestid"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func UnaryServerInterceptor(ops ...Option) grpc.UnaryServerInterceptor {
	o := initOption(ops...)
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		rid := mdutil.ValueFromContext(ctx, o.Header)
		if len(rid) == 0 {
			rid = o.Generator()
		}
		outMD := metadata.New(map[string]string{
			o.Header: rid,
		})
		resp, err = handler(metadata.NewOutgoingContext(ctx, outMD), req)
		return
	}
}

func StreamServerInterceptor(ops ...Option) grpc.StreamServerInterceptor {
	o := initOption(ops...)
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		rid := mdutil.ValueFromContext(ss.Context(), o.Header)
		if len(rid) == 0 {
			rid = o.Generator()
		}
		outMD := metadata.New(map[string]string{
			o.Header: rid,
		})
		wsCtx := requestid.NewContext(ss.Context(), rid)
		ws := util.WrapServerStream(ss)
		ws.WrappedContext = metadata.NewOutgoingContext(wsCtx, outMD)
		ss = ws
		return handler(srv, ss)
	}
}
