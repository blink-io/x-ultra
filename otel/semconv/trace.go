package semconv

import (
	"go.opentelemetry.io/otel/attribute"
)

const (
	DBAccessorKey = attribute.Key("db.accessor")

	DBHostPortKey = attribute.Key("db.host_port")

	DBDriverName = attribute.Key("db.driver_name")

	DBRegionKey = attribute.Key("db.region")

	DBClientName = attribute.Key("db.client_name")

	DBTimeZone = attribute.Key("db.time_zone")

	DBProvider = attribute.Key("db.provider")

	DBProviderAlicloud = attribute.Key("db.provider.alicloud")

	DBProviderAWS = attribute.Key("db.provider.aws")

	DBProviderAzure = attribute.Key("db.provider.azure")

	DBProviderGCP = attribute.Key("db.provider.gcp")
)

const (
	ApplicationNameKey = attribute.Key("application.name")

	ApplicationVersionKey = attribute.Key("application.version")
)
