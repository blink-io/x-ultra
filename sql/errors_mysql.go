package sql

import (
	"github.com/blink-io/x/cast"
	"github.com/go-sql-driver/mysql"
)

func mysqlStateErr(e *mysql.MySQLError) *StateError {
	err := &StateError{
		cause:   e,
		code:    cast.ToString(e.Number),
		message: e.Message,
	}
	return err
}
