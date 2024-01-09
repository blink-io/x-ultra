package db

import (
	"context"
	"net"
	"time"

	"github.com/blink-io/x/cast"
	mysqlparams "github.com/blink-io/x/mysql/params"
	"github.com/blink-io/x/sql"
	xsql "github.com/blink-io/x/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/schema"
)

func init() {
	dialectors[xsql.DialectMySQL] = NewMySQLDialect
}

func NewMySQLDialect(ctx context.Context, ops ...DialectOption) schema.Dialect {
	return mysqldialect.New()
}

func MySQLDSN(ctx context.Context, c *sql.Config) (string, error) {
	cc := ToMySQLConfig(c)
	dsn := cc.FormatDSN()
	return dsn, nil
}

func ToMySQLConfig(c *sql.Config) *mysql.Config {
	network := c.Network
	name := c.Name
	host := c.Host
	port := c.Port
	user := c.User
	password := c.Password
	dialTimeout := c.DialTimeout
	tlsConfig := c.TLSConfig
	loc := c.Loc
	collation := c.Collation
	params := c.Params

	if loc == nil {
		loc = time.Local
	}
	if params == nil {
		params = make(map[string]string)
	}
	if len(c.ClientName) > 0 {
		params[mysqlparams.ProgramName] = c.ClientName
	}
	if len(c.Collation) > 0 {
		params[mysqlparams.Collation] = c.Collation
	}

	// Restful TLS Params
	cc := mysql.NewConfig()
	// Put the local timezone because the default value is UTC
	cc.Loc = loc
	cc.ParseTime = true
	cc.Net = network
	cc.DBName = name
	cc.User = user
	cc.Passwd = password
	if dialTimeout > 0 {
		cc.Timeout = dialTimeout
	}
	// TODO Do we need to check them?
	cc.Params = handleMySQLParams(params)
	cc.Collation = collation
	if network == "tcp" {
		cc.Addr = net.JoinHostPort(host, cast.ToString(port))
	} else {
		// Otherwise, addr is Unix domain sockets
		cc.Addr = host
	}
	if tlsConfig != nil {
		keyName := mysqlTLSKeyName(name)
		_ = mysql.RegisterTLSConfig(keyName, tlsConfig)
		cc.TLSConfig = keyName
	}
	return cc
}

func mysqlTLSKeyName(name string) string {
	return xsql.DialectMySQL + "_" + name
}

func handleMySQLParams(params map[string]string) map[string]string {
	newParams := make(map[string]string)
	for k, v := range params {
		newParams[k] = v
	}
	return newParams
}
