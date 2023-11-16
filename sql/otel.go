package sql

import (
	"database/sql"
	"database/sql/driver"

	xotelsql "github.com/XSAM/otelsql"
	"github.com/uptrace/opentelemetry-go-extra/otelsql"
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

const (
	// DBAccessMethodKey represents ORM, DAL, etc.
	DBAccessMethodKey = attribute.Key("db.access_method")
)

type otelOptions struct {
	dbName         string
	dbSystem       string
	dbAccessMethod string //dbDam, Database Access Method
	reportDBStats  bool
}

type otelOption func(*otelOptions)

func (oo *otelOptions) forOTelSql() []otelsql.Option {
	ops := make([]otelsql.Option, 0)
	attrs := make([]attribute.KeyValue, 0)
	attrs = oo.appendDBInfo(attrs)
	ops = append(ops, otelsql.WithAttributes(attrs...))
	return ops
}

func (oo *otelOptions) forXOTelSql() []xotelsql.Option {
	ops := make([]xotelsql.Option, 0)
	attrs := make([]attribute.KeyValue, 0)
	attrs = oo.appendDBInfo(attrs)
	ops = append(ops, xotelsql.WithAttributes(attrs...))
	return ops
}

func (oo *otelOptions) appendDBInfo(attrs []attribute.KeyValue) []attribute.KeyValue {
	if len(oo.dbName) > 0 {
		attrs = append(attrs, semconv.DBNameKey.String(oo.dbName))
	}
	if len(oo.dbSystem) > 0 {
		attrs = append(attrs, semconv.DBSystemKey.String(oo.dbSystem))
	}
	if len(oo.dbAccessMethod) > 0 {
		attrs = append(attrs, DBAccessMethodKey.String(oo.dbAccessMethod))
	}
	return attrs
}

func applyOtelOptions(ops ...otelOption) *otelOptions {
	oop := new(otelOptions)
	for _, oo := range ops {
		oo(oop)
	}
	return oop
}

func OTelDBName(dbName string) otelOption {
	return func(o *otelOptions) {
		o.dbName = dbName
	}
}

func OTelDBSystem(dbName string) otelOption {
	return func(o *otelOptions) {
		o.dbName = dbName
	}
}

func OTelDBAccessMethod(dbAccessMethod string) otelOption {
	return func(o *otelOptions) {
		o.dbAccessMethod = dbAccessMethod
	}
}

func OTelReportDBStats() otelOption {
	return func(o *otelOptions) {
		o.reportDBStats = true
	}
}

func otelOpenDB(cc driver.Connector, ops ...otelOption) *sql.DB {
	oop := applyOtelOptions(ops...)
	db := otelsql.OpenDB(cc, oop.forOTelSql()...)
	if oop.reportDBStats {
		otelsql.ReportDBStatsMetrics(db)
	}
	return db
}

func xotelOpenDB(cc driver.Connector, ops ...otelOption) *sql.DB {
	oop := applyOtelOptions(ops...)
	db := xotelsql.OpenDB(cc, oop.forXOTelSql()...)
	if oop.reportDBStats {
		xotelsql.RegisterDBStatsMetrics(db)
	}
	return db
}

func sqlOpenDB(cc driver.Connector, ops ...otelOption) *sql.DB {
	return sql.OpenDB(cc)
}
