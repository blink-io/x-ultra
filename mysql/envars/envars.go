package envars

// Source: https://dev.mysql.com/doc/refman/8.2/en/environment-variables.html
const (
	// MYSQL_HOST specifies the default host name used by the mysql command-line client.
	MYSQL_HOST = "MYSQL_HOST"

	// MYSQL_DEBUG enables debug trace options when debugging.
	MYSQL_DEBUG = "MYSQL_DEBUG"

	// MYSQL_PWD specifies the default password when connecting to mysqld. Using this is insecure.
	// Use of MYSQL_PWD to specify a MySQL password must be considered extremely insecure and should not be used.
	// MYSQL_PWD is deprecated as of MySQL 8.2; expect it to be removed in a future version of MySQL.
	MYSQL_PWD = "MYSQL_PWD"

	// MYSQL_TCP_PORT specifies the default TCP/IP port number.
	MYSQL_TCP_PORT = "MYSQL_TCP_PORT"

	// MYSQL_UNIX_PORT specifies the default Unix socket file name; used for connections to localhost.
	MYSQL_UNIX_PORT = "MYSQL_UNIX_PORT"

	// TZ defines this should be set to your local time zone.
	// See https://dev.mysql.com/doc/refman/8.2/en/timezone-problems.html
	TZ = "TZ"

	// USER specifies the default user name on Windows when connecting to mysqld.
	USER = "USER"
)
