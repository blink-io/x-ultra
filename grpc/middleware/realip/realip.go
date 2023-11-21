package realip

import (
	"context"

	"github.com/blink-io/x/grpc/util"
	"github.com/blink-io/x/realip"
	"google.golang.org/grpc"
)

func UnaryClientInterceptor(ops ...Option) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		o := initOption(ops...)
		if ip := o.GetFromGRPC(ctx); len(ip) > 0 {
			ctx = realip.NewContext(ctx, ip)
		}
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func StreamClientInterceptor(ops ...Option) grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		o := initOption(ops...)
		if ip := o.GetFromGRPC(ctx); len(ip) > 0 {
			ctx = realip.NewContext(ctx, ip)
		}
		return streamer(ctx, desc, cc, method, opts...)
	}
}

func UnaryServerInterceptor(ops ...Option) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		o := initOption(ops...)
		if ip := o.GetFromGRPC(ctx); len(ip) > 0 {
			ctx = realip.NewContext(ctx, ip)
		}
		resp, err = handler(ctx, req)
		return
	}
}

func StreamServerInterceptor(ops ...Option) grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		o := initOption(ops...)
		ctx := ss.Context()
		if ip := o.GetFromGRPC(ctx); len(ip) > 0 {
			ws := util.WrapServerStream(ss)
			ws.WrappedContext = realip.NewContext(ctx, ip)
			ss = ws
		}
		return handler(srv, ss)
	}
}
