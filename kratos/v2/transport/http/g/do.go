package g

import (
	"context"
	"net/http"

	khttp "github.com/blink-io/x/kratos/v2/transport/http"
)

type Req = any
type Res = any

type Func[Request Req, Response Res] func(context.Context, *Request) (*Response, error)

func (h Func[Request, Response]) Do(method, operation string) khttp.HandlerFunc {
	return Do[Request, Response](method, operation, h)
}

func (h Func[Request, Response]) GET(operation string) khttp.HandlerFunc {
	return GET[Request, Response](operation, h)
}

func (h Func[Request, Response]) POST(operation string) khttp.HandlerFunc {
	return POST[Request, Response](operation, h)
}

func (h Func[Request, Response]) PUT(operation string) khttp.HandlerFunc {
	return PUT[Request, Response](operation, h)
}

func (h Func[Request, Response]) PATCH(operation string) khttp.HandlerFunc {
	return PATCH[Request, Response](operation, h)
}

func (h Func[Request, Response]) CONNECT(operation string) khttp.HandlerFunc {
	return CONNECT[Request, Response](operation, h)
}

func (h Func[Request, Response]) DELETE(operation string) khttp.HandlerFunc {
	return DELETE[Request, Response](operation, h)
}

func (h Func[Request, Response]) OPTIONS(operation string) khttp.HandlerFunc {
	return OPTIONS[Request, Response](operation, h)
}

func (h Func[Request, Response]) TRACE(operation string) khttp.HandlerFunc {
	return TRACE[Request, Response](operation, h)
}

func (h Func[Request, Response]) HEAD(operation string) khttp.HandlerFunc {
	return HEAD[Request, Response](operation, h)
}

func Do[Request Req, Response Req](
	method, operation string,
	handle func(context.Context, *Request) (*Response, error),
) khttp.HandlerFunc {
	return func(kctx khttp.Context) error {
		var in Request
		switch method {
		case http.MethodPost,
			http.MethodPut,
			// HTTP DELETE Maybe has payload
			// https://developer.mozilla.org/docs/Web/HTTP/Methods/DELETE
			http.MethodDelete,
			http.MethodPatch:
			if err := kctx.Bind(&in); err != nil {
				return err
			}
			break
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

func GET[Request Req, Response Res](
	operation string,
	handle func(context.Context, *Request) (*Response, error),
) khttp.HandlerFunc {
	return Do[Request, Response](http.MethodGet, operation, handle)
}

func POST[Request Req, Response Res](
	operation string,
	handle func(context.Context, *Request) (*Response, error)) khttp.HandlerFunc {
	return Do[Request, Response](http.MethodPost, operation, handle)
}

func PUT[Request Req, Response Res](
	operation string,
	handle func(context.Context, *Request) (*Response, error)) khttp.HandlerFunc {
	return Do[Request, Response](http.MethodPut, operation, handle)
}

func PATCH[Request Req, Response Res](
	operation string,
	handle func(context.Context, *Request) (*Response, error)) khttp.HandlerFunc {
	return Do[Request, Response](http.MethodPatch, operation, handle)
}

func DELETE[Request Req, Response Res](
	operation string,
	handle func(context.Context, *Request) (*Response, error)) khttp.HandlerFunc {
	return Do[Request, Response](http.MethodDelete, operation, handle)
}

func OPTIONS[Request Req, Response Res](
	operation string,
	handle func(context.Context, *Request) (*Response, error)) khttp.HandlerFunc {
	return Do[Request, Response](http.MethodOptions, operation, handle)
}

func CONNECT[Request Req, Response Res](
	operation string,
	handle func(context.Context, *Request) (*Response, error),
) khttp.HandlerFunc {
	return Do[Request, Response](http.MethodConnect, operation, handle)
}

func TRACE[Request Req, Response Res](
	operation string,
	handle func(context.Context, *Request) (*Response, error),
) khttp.HandlerFunc {
	return Do[Request, Response](http.MethodTrace, operation, handle)
}

func HEAD[Request Req, Response Res](
	operation string,
	handle func(context.Context, *Request) (*Response, error),
) khttp.HandlerFunc {
	return Do[Request, Response](http.MethodHead, operation, handle)
}
