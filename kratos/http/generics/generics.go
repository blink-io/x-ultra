package generics

import (
	"context"
	"net/http"

	khttp "github.com/go-kratos/kratos/v2/transport/http"
)

func Handle[Req any, Res any](
	operation string,
	dlr func(context.Context, *Req) (*Res, error),
	ops ...Option,
) khttp.HandlerFunc {
	return func(kctx khttp.Context) error {
		opts := new(options)
		for _, o := range ops {
			o(opts)
		}
		var in Req
		if opts.method == http.MethodPost ||
			opts.method == http.MethodPut ||
			// HTTP DELETE Maybe has payload
			// https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Methods/DELETE
			opts.method == http.MethodDelete ||
			opts.method == http.MethodPatch {
			if err := kctx.Bind(&in); err != nil {
				return err
			}
		}
		if err := kctx.BindQuery(&in); err != nil {
			return err
		}
		khttp.SetOperation(kctx, operation)
		mdlr := kctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return dlr(kctx, req.(*Req))
		})
		out, err := mdlr(kctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Res)
		return kctx.Result(200, reply)
	}
}