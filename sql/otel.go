package sql

import (
	"database/sql"
	"database/sql/driver"

	xsemconv "github.com/blink-io/x/otel/semconv"

	xotelsql "github.com/XSAM/otelsql"
	"github.com/uptrace/opentelemetry-go-extra/otelsql"
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

type otelOptions struct {
	dbName         string
	dbSystem       string
	dbAccessMethod string //dbDam, Database Access Method
	dbHostPort     string
	reportDBStats  bool
}

type OTtelOption func(*otelOptions)

func (oo *otelOptions) otelSqlOptions() []otelsql.Option {
	ops := make([]otelsql.Option, 0)
	attrs := oo.createDBInfoAttrs()
	ops = append(ops, otelsql.WithAttributes(attrs...))
	return ops
}

func (oo *otelOptions) xotelSqlOptions() []xotelsql.Option {
	ops := make([]xotelsql.Option, 0)
	attrs := oo.createDBInfoAttrs()
	ops = append(ops, xotelsql.WithAttributes(attrs...))
	return ops
}

func (oo *otelOptions) createDBInfoAttrs() []attribute.KeyValue {
	attrs := make([]attribute.KeyValue, 0)
	if len(oo.dbName) > 0 {
		attrs = append(attrs, semconv.DBNameKey.String(oo.dbName))
	}
	if len(oo.dbSystem) > 0 {
		attrs = append(attrs, semconv.DBSystemKey.String(oo.dbSystem))
	}
	if len(oo.dbAccessMethod) > 0 {
		attrs = append(attrs, xsemconv.DBAccessMethodKey.String(oo.dbAccessMethod))
	}
	if len(oo.dbHostPort) > 0 {
		attrs = append(attrs, xsemconv.DBHostPortKey.String(oo.dbHostPort))
	}
	return attrs
}

func applyOtelOptions(ops ...OTtelOption) *otelOptions {
	oop := new(otelOptions)
	for _, oo := range ops {
		oo(oop)
	}
	return oop
}

func OTelDBName(dbName string) OTtelOption {
	return func(o *otelOptions) {
		o.dbName = dbName
	}
}

func OTelDBSystem(dbName string) OTtelOption {
	return func(o *otelOptions) {
		o.dbName = dbName
	}
}

func OTelDBAccessMethod(dbAccessMethod string) OTtelOption {
	return func(o *otelOptions) {
		o.dbAccessMethod = dbAccessMethod
	}
}

func OTelDBHostPort(dbHostPort string) OTtelOption {
	return func(o *otelOptions) {
		o.dbHostPort = dbHostPort
	}
}

func OTelReportDBStats() OTtelOption {
	return func(o *otelOptions) {
		o.reportDBStats = true
	}
}

func otelOpenDB(cc driver.Connector, ops ...OTtelOption) *sql.DB {
	oop := applyOtelOptions(ops...)
	db := otelsql.OpenDB(cc, oop.otelSqlOptions()...)
	if oop.reportDBStats {
		otelsql.ReportDBStatsMetrics(db)
	}
	return db
}

func xotelOpenDB(cc driver.Connector, ops ...OTtelOption) *sql.DB {
	oop := applyOtelOptions(ops...)
	db := xotelsql.OpenDB(cc, oop.xotelSqlOptions()...)
	if oop.reportDBStats {
		xotelsql.RegisterDBStatsMetrics(db)
	}
	return db
}

func sqlOpenDB(cc driver.Connector, ops ...OTtelOption) *sql.DB {
	return sql.OpenDB(cc)
}
