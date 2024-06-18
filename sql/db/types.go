package db

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

	TableType = schema.Table

	QueryWithArgs = schema.QueryWithArgs

	Ident = bun.Ident

	InsertQuery = bun.InsertQuery

	DeleteQuery = bun.DeleteQuery

	UpdateQuery = bun.UpdateQuery

	SelectQuery = bun.SelectQuery
)

func In(slice interface{}) schema.QueryAppender {
	return schema.In(slice)
}

func NullZero(value interface{}) schema.QueryAppender {
	return schema.NullZero(value)
}
