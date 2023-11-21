package params

// Source: https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING
const (
	// HOST specifies name of host to connect to.
	HOST = "host"

	// PORT specifies port number to connect to at the server host, or socket file name extension for Unix-domain connections.
	PORT = "port"

	// DBNAME specifies database name.
	DBNAME = "dbname"

	// USER specifies a PostgreSQL user to connect as.
	USER = "user"

	// PASSWORD specifies password to be used if the server demands password authentication.
	PASSWORD = "password"

	// PASSFILE specifies the name of the file used to store passwords (see Section 34.16).
	// Defaults to ~/.pgpass or %APPDATA%\postgresql\pgpass.conf on Microsoft Windows.
	// passfile format: hostname:port:database:username:password
	// Source: https://www.postgresql.org/docs/current/libpq-pgpass.html
	PASSFILE = "passfile"

	// APPLICATION_NAME specifies a value for the application_name configuration parameter.
	APPLICATION_NAME = "application_name"

	// FALLBACK_APPLICATION_NAME specifies a fallback value for the application_name configuration parameter.
	FALLBACK_APPLICATION_NAME = "fallback_application_name"

	// CLIENT_ENCODING sets the client_encoding configuration parameter for this connection.
	CLIENT_ENCODING = "client_encoding"

	// SSLMODE values: disable|allow|prefer (default)|require|verify-ca|verify-full
	SSLMODE = "sslmode"

	// SSLKEY specifies the location for the secret key used for the client certificate.
	SSLKEY = "sslkey"

	// SSLCERT specifies the file name of the client SSL certificate,
	// replacing the default ~/.postgresql/postgresql.crt.
	// This parameter is ignored if an SSL connection is not made.
	SSLCERT = "sslcert"

	// SSLROOTCERT specifies the name of a file containing SSL certificate authority (CA) certificate(s).
	SSLROOTCERT = "sslrootcert"

	// SSLPASSWORD specifies the password for the secret key specified in sslkey,
	SSLPASSWORD = "sslpassword"

	// SSLCERTMODE determines whether a client certificate may be sent to the server,
	// and whether the server is required to request one. There are three modes:
	// disable|allow (default)|require
	SSLCERTMODE = "sslcertmode"

	// SSLCRL specifies the file name of the SSL server certificate revocation list (CRL).
	SSLCRL = "sslcrl"

	// SSLCRLDIR specifies the directory name of the SSL server certificate revocation list (CRL).
	SSLCRLDIR = "sslcrldir"

	// SERVICE specifies service name to use for additional parameters.
	SERVICE = "service"

	// CONNECT_TIMEOUT specifies maximum time to wait while connecting, in seconds (write as a decimal integer, e.g., 10).
	// Zero, negative, or not specified means wait indefinitely.
	CONNECT_TIMEOUT = "connect_timeout"

	LOAD_BALANCE_HOSTS = "" //optional values: disable(default)|random
)
