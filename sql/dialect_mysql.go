package sql

import (
	"context"
	"database/sql/driver"
	"net"
	"time"

	"github.com/blink-io/x/cast"
	mysqlparams "github.com/blink-io/x/mysql/params"
	"github.com/go-sql-driver/mysql"
)

var compatibleMySQLDialects = []string{
	DialectMySQL,
	"mysql5",
	"mysql8",
}

func init() {
	d := DialectMySQL
	//drivers[d] = GetMySQLDriver
	//dsners[d] = GetMySQLDSN
	connectors[d] = GetMySQLConnector
}

type MySQLOptions struct {
}

func ValidateMySQLConfig(c *Config) error {
	return nil
}

func GetMySQLDSN(dialect string) (Dsner, error) {
	if !IsCompatibleMySQLDialect(dialect) {
		return nil, ErrUnsupportedDialect
	}
	return func(ctx context.Context, c *Config) (string, error) {
		cc := ToMySQLConfig(c)
		dsn := cc.FormatDSN()
		return dsn, nil
	}, nil
}

func GetMySQLDriver(dialect string) (driver.Driver, error) {
	if IsCompatibleMySQLDialect(dialect) {
		return &mysql.MySQLDriver{}, nil
	}
	return nil, ErrUnsupportedDriver
}

func GetMySQLConnector(ctx context.Context, c *Config) (driver.Connector, error) {
	cc := ToMySQLConfig(c)
	dsn := cc.FormatDSN()
	drv := wrapDriverHooks(getRawPostgresDriver(), c.DriverHooks...)
	return &dsnConnector{dsn: dsn, driver: drv}, nil
}

func ToMySQLConfig(c *Config) *mysql.Config {
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
	// Force to parse to time.Time
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

func IsCompatibleMySQLDialect(dialect string) bool {
	return isCompatibleDialectIn(dialect, compatibleMySQLDialects)
}

func AdditionsToMySQLOptions(adds map[string]string) *MySQLOptions {
	opts := new(MySQLOptions)
	return opts
}

func getRawMySQLDriver() driver.Driver {
	return &mysql.MySQLDriver{}
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
