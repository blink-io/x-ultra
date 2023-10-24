//go:build !otel

package sql

import (
	"database/sql"
)

var openDB = sql.OpenDB
