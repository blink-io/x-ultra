package connect

type HandlerWrapper[T any] struct {
	Prefix  string
	Handler T
}

func NewHandlerWrapper[H any](prefix string, h H) *HandlerWrapper[H] {
	return &HandlerWrapper[H]{
		Prefix:  prefix,
		Handler: h,
	}
}
