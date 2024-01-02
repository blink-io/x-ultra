package http

import (
	"github.com/blink-io/x/kratos/v2/transport/httpbase"
)

type (
	Client = httpbase.Client

	ClientOption = httpbase.ClientOption

	Context = httpbase.Context

	RouteInfo = httpbase.RouteInfo
)

var (
	NewClient = httpbase.NewClient

	WithTLSConfig = httpbase.WithTLSConfig

	WithEndpoint = httpbase.WithEndpoint

	WithTransport = httpbase.WithTransport
)
