package sql

import (
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
	dialectFuncs[dn] = func() schema.Dialect {
		return mysqldialect.New()
	}
	dsnFuncs[dn] = MySQLDSN
}

func MySQLDSN(o *Options) string {
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
	// Set the local timezone because the default value is UTC
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
		// Driver Name is valid ant will not throw an error
		_ = mysql.RegisterTLSConfig(DialectMySQL, tlsConfig)
		cc.TLSConfig = DialectMySQL
	}
	dsn := cc.FormatDSN()
	return dsn
}

func handleMySQLParams(params map[string]string) map[string]string {
	newParams := make(map[string]string)
	for k, v := range params {
		newParams[k] = v
	}
	return newParams
}
