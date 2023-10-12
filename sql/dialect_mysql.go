//go:build mysql

package sql

import (
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

	slog.Info("MySQL is enabled")
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
	options := o.Options
	loc := o.Loc

	// Restful TLS Options
	cc := mysql.NewConfig()
	// Set the local timezone because the default value is UTC
	cc.Loc = loc
	cc.ParseTime = true
	cc.Net = network
	cc.DBName = name
	cc.User = user
	cc.Timeout = dialTimeout
	cc.Passwd = password
	// TODO Do we need to check them?
	cc.Params = options
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
}
