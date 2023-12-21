package sql

import (
	"context"
	"net"

	"github.com/blink-io/x/cast"
	mysqlparams "github.com/blink-io/x/mysql/params"

	"github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/schema"
)

const (
	DialectMySQL = "mysql"
)

func init() {
	dn := DialectMySQL
	drivers[dn] = &mysql.MySQLDriver{}
	dialectCreators[dn] = func(ctx context.Context, ops ...DOption) schema.Dialect {
		return mysqldialect.New()
	}
	dsnCreators[dn] = MySQLDSN
}

func MySQLDSN(o *Options) (string, error) {
	cc := ToMySQLConfig(o)
	dsn := cc.FormatDSN()
	return dsn, nil
}

func ToMySQLConfig(o *Options) *mysql.Config {
	network := o.Network
	name := o.Name
	host := o.Host
	port := o.Port
	user := o.User
	password := o.Password
	dialTimeout := o.DialTimeout
	tlsConfig := o.TLSConfig
	loc := o.Loc
	collation := o.Collation
	params := o.Params
	if params == nil {
		params = make(map[string]string)
	}
	if len(o.ClientName) > 0 {
		params[mysqlparams.ProgramName] = o.ClientName
	}
	if len(o.Collation) > 0 {
		params[mysqlparams.Collation] = o.Collation
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
	return DialectMySQL + "_" + name
}

func handleMySQLParams(params map[string]string) map[string]string {
	newParams := make(map[string]string)
	for k, v := range params {
		newParams[k] = v
	}
	return newParams
}
