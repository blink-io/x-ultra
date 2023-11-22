package params

// Source: https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING
const (
	// Host specifies name of host to connect to.
	Host = "host"

	// Port specifies port number to connect to at the server host, or socket file name extension for Unix-domain connections.
	Port = "port"

	// DBName specifies database name.
	DBName = "dbname"

	// User specifies a PostgreSQL user to connect as.
	User = "user"

	// Password specifies password to be used if the server demands password authentication.
	Password = "password"

	// Passfile specifies the name of the file used to store passwords (see Section 34.16).
	// Defaults to ~/.pgpass or %APPDATA%\postgresql\pgpass.conf on Microsoft Windows.
	// passfile format: hostname:port:database:username:password
	// Source: https://www.postgresql.org/docs/current/libpq-pgpass.html
	Passfile = "passfile"

	// ApplicationName specifies a value for the application_name configuration parameter.
	ApplicationName = "application_name"

	// FallbackApplicationName specifies a fallback value for the application_name configuration parameter.
	FallbackApplicationName = "fallback_application_name"

	// ClientEncoding sets the client_encoding configuration parameter for this connection.
	ClientEncoding = "client_encoding"

	// SSLMode values: disable|allow|prefer (default)|require|verify-ca|verify-full
	SSLMode = "sslmode"

	// SSLKey specifies the location for the secret key used for the client certificate.
	SSLKey = "sslkey"

	// SSLCert specifies the file name of the client SSL certificate,
	// replacing the default ~/.postgresql/postgresql.crt.
	// This parameter is ignored if an SSL connection is not made.
	SSLCert = "sslcert"

	// SSLRootCert specifies the name of a file containing SSL certificate authority (CA) certificate(s).
	SSLRootCert = "sslrootcert"

	// SSLPassword specifies the password for the secret key specified in sslkey,
	SSLPassword = "sslpassword"

	// SSLCertMode determines whether a client certificate may be sent to the server,
	// and whether the server is required to request one. There are three modes:
	// disable|allow (default)|require
	SSLCertMode = "sslcertmode"

	// SSLCRL specifies the file name of the SSL server certificate revocation list (CRL).
	SSLCRL = "sslcrl"

	// SSLCRLDir specifies the directory name of the SSL server certificate revocation list (CRL).
	SSLCRLDir = "sslcrldir"

	// Service specifies service name to use for additional parameters.
	Service = "service"

	// ConnectTimeout specifies maximum time to wait while connecting, in seconds (write as a decimal integer, e.g., 10).
	// Zero, negative, or not specified means wait indefinitely.
	ConnectTimeout = "connect_timeout"

	LoadBalanceHosts = "" //optional values: disable(default)|random
)
