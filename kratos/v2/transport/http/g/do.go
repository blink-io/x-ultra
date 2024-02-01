package g

import (
	"context"
	"net/http"

	khttp "github.com/blink-io/x/kratos/v2/transport/http"
)

type Req = any
type Res = any

type Func[Request Req, Response Res] func(context.Context, *Request) (*Response, error)

func (h Func[Request, Response]) Do(method, operation string, ops ...DoOption) khttp.HandlerFunc {
	return Do[Request, Response](method, operation, h, ops...)
}

func (h Func[Request, Response]) GET(operation string, ops ...DoOption) khttp.HandlerFunc {
	return GET[Request, Response](operation, h, ops...)
}

func (h Func[Request, Response]) POST(operation string, ops ...DoOption) khttp.HandlerFunc {
	return POST[Request, Response](operation, h, ops...)
}

func (h Func[Request, Response]) PUT(operation string, ops ...DoOption) khttp.HandlerFunc {
	return PUT[Request, Response](operation, h, ops...)
}

func (h Func[Request, Response]) PATCH(operation string, ops ...DoOption) khttp.HandlerFunc {
	return PATCH[Request, Response](operation, h, ops...)
}

func (h Func[Request, Response]) CONNECT(operation string, ops ...DoOption) khttp.HandlerFunc {
	return CONNECT[Request, Response](operation, h, ops...)
}

func (h Func[Request, Response]) DELETE(operation string, ops ...DoOption) khttp.HandlerFunc {
	return DELETE[Request, Response](operation, h, ops...)
}

func (h Func[Request, Response]) OPTIONS(operation string, ops ...DoOption) khttp.HandlerFunc {
	return OPTIONS[Request, Response](operation, h, ops...)
}

func (h Func[Request, Response]) TRACE(operation string, ops ...DoOption) khttp.HandlerFunc {
	return TRACE[Request, Response](operation, h, ops...)
}

func (h Func[Request, Response]) HEAD(operation string, ops ...DoOption) khttp.HandlerFunc {
	return HEAD[Request, Response](operation, h, ops...)
}

func Do[Request Req, Response Req](
	method, operation string,
	handle func(context.Context, *Request) (*Response, error),
	ops ...DoOption,
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
		if err := kctx.BindVars(&in); err != nil {
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

		opts := applyOptions(ops...)
		for _, i := range opts.ahs {
			kctx.Header().Add(i.key, i.value)
		}
		for k, v := range opts.shs {
			kctx.Header().Set(k, v)
		}

		reply := out.(*Response)
		return kctx.Result(opts.statusCode, reply)
	}
}

func GET[Request Req, Response Res](
	operation string,
	handle func(context.Context, *Request) (*Response, error),
	ops ...DoOption,
) khttp.HandlerFunc {
	return Do[Request, Response](http.MethodGet, operation, handle, ops...)
}

func POST[Request Req, Response Res](
	operation string,
	handle func(context.Context, *Request) (*Response, error),
	ops ...DoOption,
) khttp.HandlerFunc {
	return Do[Request, Response](http.MethodPost, operation, handle, ops...)
}

func PUT[Request Req, Response Res](
	operation string,
	handle func(context.Context, *Request) (*Response, error),
	ops ...DoOption,
) khttp.HandlerFunc {
	return Do[Request, Response](http.MethodPut, operation, handle, ops...)
}

func PATCH[Request Req, Response Res](
	operation string,
	handle func(context.Context, *Request) (*Response, error),
	ops ...DoOption,
) khttp.HandlerFunc {
	return Do[Request, Response](http.MethodPatch, operation, handle, ops...)
}

func DELETE[Request Req, Response Res](
	operation string,
	handle func(context.Context, *Request) (*Response, error),
	ops ...DoOption,
) khttp.HandlerFunc {
	return Do[Request, Response](http.MethodDelete, operation, handle, ops...)
}

func OPTIONS[Request Req, Response Res](
	operation string,
	handle func(context.Context, *Request) (*Response, error),
	ops ...DoOption,
) khttp.HandlerFunc {
	return Do[Request, Response](http.MethodOptions, operation, handle, ops...)
}

func CONNECT[Request Req, Response Res](
	operation string,
	handle func(context.Context, *Request) (*Response, error),
	ops ...DoOption,
) khttp.HandlerFunc {
	return Do[Request, Response](http.MethodConnect, operation, handle, ops...)
}

func TRACE[Request Req, Response Res](
	operation string,
	handle func(context.Context, *Request) (*Response, error),
	ops ...DoOption,
) khttp.HandlerFunc {
	return Do[Request, Response](http.MethodTrace, operation, handle, ops...)
}

func HEAD[Request Req, Response Res](
	operation string,
	handle func(context.Context, *Request) (*Response, error),
	ops ...DoOption,
) khttp.HandlerFunc {
	return Do[Request, Response](http.MethodHead, operation, handle, ops...)
}
