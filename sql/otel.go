//go:build otel

package sql

import (
	"github.com/uptrace/opentelemetry-go-extra/otelsql"
)

var openDB = otelsql.OpenDB
