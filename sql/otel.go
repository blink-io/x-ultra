package sql

import (
	"database/sql"
	"database/sql/driver"

	xsemconv "github.com/blink-io/x/otel/semconv"

	xotelsql "github.com/XSAM/otelsql"
	"github.com/uptrace/opentelemetry-go-extra/otelsql"
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
)

type otelOptions struct {
	dbName        string
	dbSystem      string
	dbAccessor    string //dbDam, Database Access Method
	dbHostPort    string
	extraAttrs    []attribute.KeyValue
	reportDBStats bool
}

type OTelOption func(*otelOptions)

func (oo *otelOptions) otelSqlOptions() []otelsql.Option {
	attrs := oo.createAttrs()

	ops := make([]otelsql.Option, 0)
	ops = append(ops, otelsql.WithAttributes(attrs...))
	return ops
}

func (oo *otelOptions) xotelSqlOptions() []xotelsql.Option {
	attrs := oo.createAttrs()

	ops := make([]xotelsql.Option, 0)
	ops = append(ops, xotelsql.WithAttributes(attrs...))
	return ops
}

func (oo *otelOptions) createAttrs() []attribute.KeyValue {
	attrs := make([]attribute.KeyValue, 0)
	if len(oo.extraAttrs) > 0 {
		attrs = append(attrs, oo.extraAttrs...)
	}
	attrs = append(attrs, oo.createDBInfoAttrs()...)
	// TODO Remove duplicated attrs
	return attrs
}

func (oo *otelOptions) createDBInfoAttrs() []attribute.KeyValue {
	attrs := make([]attribute.KeyValue, 0)
	if len(oo.dbName) > 0 {
		attrs = append(attrs, semconv.DBNameKey.String(oo.dbName))
	}
	if len(oo.dbSystem) > 0 {
		attrs = append(attrs, semconv.DBSystemKey.String(oo.dbSystem))
	}
	if len(oo.dbAccessor) > 0 {
		attrs = append(attrs, xsemconv.DBAccessorKey.String(oo.dbAccessor))
	}
	if len(oo.dbHostPort) > 0 {
		attrs = append(attrs, xsemconv.DBHostPortKey.String(oo.dbHostPort))
	}
	return attrs
}

func applyOtelOptions(ops ...OTelOption) *otelOptions {
	oop := new(otelOptions)
	for _, oo := range ops {
		oo(oop)
	}
	return oop
}

func OTelDBName(dbName string) OTelOption {
	return func(o *otelOptions) {
		o.dbName = dbName
	}
}

func OTelDBSystem(dbName string) OTelOption {
	return func(o *otelOptions) {
		o.dbName = dbName
	}
}

func OTelDBAccessor(dbAccessor string) OTelOption {
	return func(o *otelOptions) {
		o.dbAccessor = dbAccessor
	}
}

func OTelDBHostPort(dbHostPort string) OTelOption {
	return func(o *otelOptions) {
		o.dbHostPort = dbHostPort
	}
}

func OTelReportDBStats() OTelOption {
	return func(o *otelOptions) {
		o.reportDBStats = true
	}
}

func OTelAttrs(attrs ...attribute.KeyValue) OTelOption {
	return func(o *otelOptions) {
		o.extraAttrs = attrs
	}
}

func otelOpenDB(cc driver.Connector, ops ...OTelOption) *sql.DB {
	oop := applyOtelOptions(ops...)
	db := otelsql.OpenDB(cc, oop.otelSqlOptions()...)
	if oop.reportDBStats {
		otelsql.ReportDBStatsMetrics(db)
	}
	return db
}

func xotelOpenDB(cc driver.Connector, ops ...OTelOption) *sql.DB {
	oop := applyOtelOptions(ops...)
	db := xotelsql.OpenDB(cc, oop.xotelSqlOptions()...)
	if oop.reportDBStats {
		_ = xotelsql.RegisterDBStatsMetrics(db)
	}
	return db
}

func otelWrapper(f func(driver.Connector) *sql.DB) func(driver.Connector, ...OTelOption) *sql.DB {
	return func(cc driver.Connector, ops ...OTelOption) *sql.DB {
		return f(cc)
	}
}
