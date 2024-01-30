package sql

type IDBExt interface {
	WithSqlDB

	WithDBInfo

	HealthChecker
}
