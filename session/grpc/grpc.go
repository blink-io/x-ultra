package grpc

import (
	"context"
	"strings"

	"github.com/blink-io/x/grpc/mdutil"
	"github.com/blink-io/x/grpc/util"
	"github.com/blink-io/x/session"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const DefaultHeader = "X-Auth-Token"

type SessionHandler struct {
	sm *session.Manager

	header string
}

func NewSessionHandler(ops ...Option) *SessionHandler {
	sh := &SessionHandler{
		sm: session.NewManager(),
	}
	for _, o := range ops {
		o(sh)
	}
	return sh
}

func (sh *SessionHandler) UnaryServerInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	sm := sh.sm
	header := strings.ToLower(sh.header)
	token := mdutil.SingleValueFromContext(ctx, header)

	ctx, err := sm.Load(ctx, token)
	if err != nil {
		return nil, err
	}

	ctx = session.NewContext(ctx, sh.sm)
	res, err := handler(ctx, req)

	if err := sh.commitAndWriteSession(ctx); err != nil {
		return nil, err
	}

	return res, err
}

func (sh *SessionHandler) StreamServerInterceptor(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	ctx := session.NewContext(ss.Context(), sh.sm)
	wss := util.WrapServerStream(ss)
	wss.WrappedContext = ctx

	sm := sh.sm
	header := strings.ToLower(sh.header)
	token := mdutil.SingleValueFromContext(ctx, header)

	ctx, err := sm.Load(ctx, token)
	if err != nil {
		return err
	}

	if err := handler(srv, ss); err != nil {
		return err
	}

	if err := sh.commitAndWriteSession(ctx); err != nil {
		return err
	}

	return nil
}

func (sh *SessionHandler) commitAndWriteSession(ctx context.Context) error {
	header := sh.header
	sm := sh.sm
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		md = metadata.Pairs()
	}

	switch sm.Status(ctx) {
	case session.Modified:
		token, _, err := sm.Commit(ctx)
		if err != nil {
			return err
		}
		md.Set(header, token)
	case session.Destroyed:
		md.Delete(header)
	}

	return grpc.SendHeader(ctx, md)
}
