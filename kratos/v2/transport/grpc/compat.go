package grpc

import (
	kgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	"google.golang.org/grpc"
)

type (
	ServiceRegistrar = grpc.ServiceRegistrar

	ServerOption = kgrpc.ServerOption
)

// For Server
var (
	Address = kgrpc.Address

	Network = kgrpc.Network

	Endpoint = kgrpc.Endpoint

	Middleware = kgrpc.Middleware

	CustomHealth = kgrpc.CustomHealth

	Timeout = kgrpc.Timeout

	Listener = kgrpc.Listener

	Options = kgrpc.Options

	UnaryInterceptor = kgrpc.UnaryInterceptor

	StreamInterceptor = kgrpc.StreamInterceptor

	TLSConfig = kgrpc.TLSConfig
)

// For Client
var (
	DialInsecure = kgrpc.DialInsecure

	WithEndpoint = kgrpc.WithEndpoint

	WithTimeout = kgrpc.WithTimeout

	WithMiddleware = kgrpc.WithMiddleware

	WithTLSConfig = kgrpc.WithTLSConfig

	WithUnaryInterceptor = kgrpc.WithUnaryInterceptor

	WithStreamInterceptor = kgrpc.WithStreamInterceptor

	WithOptions = kgrpc.WithOptions

	WithNodeFilter = kgrpc.WithNodeFilter

	WithHealthCheck = kgrpc.WithHealthCheck

	WithDiscovery = kgrpc.WithDiscovery

	WithSubset = kgrpc.WithSubset

	WithPrintDiscoveryDebugLog = kgrpc.WithPrintDiscoveryDebugLog
)
