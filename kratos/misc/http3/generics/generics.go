package generics

import (
	"context"
	"net/http"

	khttp3 "github.com/blink-io/x/kratos/transport/http3"
)

func Handle[Req any, Res any](
	operation string,
	handle func(context.Context, *Req) (*Res, error),
	ops ...Option,
) khttp3.HandlerFunc {
	return func(kctx khttp3.Context) error {
		opts := applyOptions(ops...)
		var in Req
		switch opts.method {
		case http.MethodPost,
			http.MethodPut,
			// HTTP DELETE Maybe has payload
			// https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Methods/DELETE
			http.MethodDelete,
			http.MethodPatch:
			if err := kctx.Bind(&in); err != nil {
				return err
			}
		default:
		}

		if err := kctx.BindQuery(&in); err != nil {
			return err
		}
		khttp3.SetOperation(kctx, operation)
		mHandle := kctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return handle(kctx, req.(*Req))
		})
		out, err := mHandle(kctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Res)
		return kctx.Result(http.StatusOK, reply)
	}
}
