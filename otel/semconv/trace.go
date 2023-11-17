package semconv

import (
	"go.opentelemetry.io/otel/attribute"
)

const (
	DBAccessMethodKey = attribute.Key("db.access_method")

	DBHostPortKey = attribute.Key("db.host_port")
)
