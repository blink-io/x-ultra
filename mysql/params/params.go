package params

// Source: https://dev.mysql.com/doc/refman/8.2/en/connecting-using-uri-or-key-value-pairs.html

const (
	// Host specifies the host on which the server instance is running.
	Host = "host"

	// Port specifies the TCP/IP network port on which the target MySQL server is listening for connections.
	Port = "port"

	// Socket specifies the path to a Unix socket file or the name of a Windows named pipe.
	Socket = "socket"

	// Schema specifies the default database for the connection.
	Schema = "schema"

	// User specifies the MySQL user account to provide for the authentication process.
	User = "user"

	// Password specifies the password to use for the authentication process.
	Password = "password"

	// SSLMode desired security state for the connection. The following modes are permissible:
	// DISABLED | PREFERRED | REQUIRED | VERIFY_CA | VERIFY_IDENTITY
	SSLMode = "ssl-mode"

	// SSLCA specifies the path to the X.509 certificate authority file in PEM format.
	SSLCA = "ssl-ca"

	// SSLCAPath specifies the path to the directory that contains the X.509 certificates authority files in PEM format.
	SSLCAPath = "ssl-ca-path"

	// SSLCert specifies the path to the X.509 certificate file in PEM format.
	SSLCert = "ssl-cert"

	// SSLCRL specifies the path to the file that contains certificate revocation lists in PEM format.
	SSLCRL = "ssl-crl"

	// SSLCrlpath specifies the path to the directory that contains certificate revocation-list files in PEM format.
	SSLCrlpath = "ssl-crlpath"

	// SSLKey specifies the path to the X.509 key file in PEM format.
	SSLKey = "ssl-crl"

	// TLSVersion specifies the TLS protocols permitted for classic MySQL protocol encrypted connections.
	TLSVersion = "tls-version"

	// AutoMethod specifies the authentication method to use for the connection.
	// The default is AUTO
	AutoMethod = "auth-method"

	// ConnectTimeout defines an integer value used to configure the number of seconds that clients,
	// such as MySQL Shell, wait until they stop trying to connect to an unresponsive MySQL server.
	ConnectTimeout = "connect_timeout"

	Compression = "compression"

	ProgramName = "program_name"

	Collation = "collation"
)
