package i18n

import (
	"context"

	"github.com/blink-io/x/grpc/mdutil"
	"github.com/blink-io/x/grpc/util"
	"github.com/blink-io/x/i18n"

	"google.golang.org/grpc"
)

func UnaryServerInterceptor(ops ...Option) grpc.UnaryServerInterceptor {
	o := initOption(ops...)
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if val := mdutil.SingleValueFromContext(ctx, o.header); len(val) > 0 {
			ctx = i18n.NewContext(ctx, val)
		}
		resp, err = handler(ctx, req)
		return
	}
}

func StreamServerInterceptor(ops ...Option) grpc.StreamServerInterceptor {
	o := initOption(ops...)
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		ctx := ss.Context()
		if val := mdutil.SingleValueFromContext(ctx, o.header); len(val) > 0 {
			ws := util.WrapServerStream(ss)
			ws.WrappedContext = i18n.NewContext(ctx, val)
			ss = ws
		}
		return handler(srv, ss)
	}
}
