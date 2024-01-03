package generics

import (
	"context"
	"net/http"

	khttp "github.com/blink-io/x/kratos/v2/transport/http"
)

func Handle[Request any, Response any](
	operation string,
	handle func(context.Context, *Request) (*Response, error),
	ops ...Option,
) khttp.HandlerFunc {
	return func(kctx khttp.Context) error {
		opts := applyOptions(ops...)
		var in Request
		switch opts.method {
		case http.MethodPost,
			http.MethodPut,
			// HTTP DELETE Maybe has payload
			// https://developer.mozilla.org/docs/Web/HTTP/Methods/DELETE
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
		khttp.SetOperation(kctx, operation)
		mwHandle := kctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return handle(kctx, req.(*Request))
		})
		out, err := mwHandle(kctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Response)
		return kctx.Result(http.StatusOK, reply)
	}
}
