package bun

import (
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
)

type (
	rdb = bun.DB

	RawIConn = bun.IConn

	RawConn = bun.Conn

	RawDB = bun.DB

	RawIDB = bun.IDB

	RawTx = bun.Tx

	BaseModel = bun.BaseModel

	TableType = schema.Table

	QueryWithArgs = schema.QueryWithArgs

	Ident = schema.Ident

	Name = schema.Name

	Safe = schema.Safe

	InsertQuery = bun.InsertQuery

	DeleteQuery = bun.DeleteQuery

	UpdateQuery = bun.UpdateQuery

	SelectQuery = bun.SelectQuery

	QueryBuilder = bun.QueryBuilder
)

func In(slice interface{}) schema.QueryAppender {
	return schema.In(slice)
}

func NullZero(value interface{}) schema.QueryAppender {
	return schema.NullZero(value)
}

func SafeQuery(query string, args []any) schema.QueryWithArgs {
	return schema.SafeQuery(query, args)
}

func SafeQueryWithSep(query string, args []any, sep string) schema.QueryWithSep {
	return schema.SafeQueryWithSep(query, args, sep)
}
