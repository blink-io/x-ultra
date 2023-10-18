package http

import (
	khttp "github.com/go-kratos/kratos/v2/transport/http"
)

type Handler interface {
	HandleHTTP(*khttp.Server)
}
