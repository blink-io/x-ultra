package sql

import (
	"database/sql/driver"
	"fmt"

	"github.com/uptrace/bun/schema"
)

type DialectFn = func() schema.Dialect

var dialectFns = make(map[string]DialectFn)
var drivers = make(map[string]driver.Driver)

func SetDialectFn(name string, dialectFn DialectFn) {
	dialectFns[name] = dialectFn
}

func GetDialectFn(name string) (DialectFn, error) {
	if fn, ok := dialectFns[name]; ok {
		return fn, nil
	}
	return nil, fmt.Errorf("unsupported dialect: %s", name)
}

func SetDriverFn(name string, driverFn Driver) {
	drivers[name] = driverFn
}

func GetDriverFn(name string) (Driver, error) {
	if fn, ok := drivers[name]; ok {
		return fn, nil
	}
	return nil, fmt.Errorf("unsupported driver: %s", name)
}
