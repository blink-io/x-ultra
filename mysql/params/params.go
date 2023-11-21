package params

// Source: https://dev.mysql.com/doc/refman/8.2/en/connecting-using-uri-or-key-value-pairs.html

const (
	// HOST specifies the host on which the server instance is running.
	HOST = "host"

	// PORT specifies the TCP/IP network port on which the target MySQL server is listening for connections.
	PORT = "port"

	// SOCKET specifies the path to a Unix socket file or the name of a Windows named pipe.
	SOCKET = ""

	// SCHEMA specifies the default database for the connection.
	SCHEMA = "schema"

	// USER specifies the MySQL user account to provide for the authentication process.
	USER = "user"

	// PASSWORD specifies the password to use for the authentication process.
	PASSWORD = "password"

	// SSL_MODE desired security state for the connection. The following modes are permissible:
	// DISABLED | PREFERRED | REQUIRED | VERIFY_CA | VERIFY_IDENTITY
	SSL_MODE = "ssl-mode"

	// SSL_CA specifies the path to the X.509 certificate authority file in PEM format.
	SSL_CA = "ssl-ca"

	// SSL_CA_PATH specifies the path to the directory that contains the X.509 certificates authority files in PEM format.
	SSL_CA_PATH = "ssl-ca-path"

	// SSL_CERT specifies the path to the X.509 certificate file in PEM format.
	SSL_CERT = "ssl-cert"

	// SSL_CRL specifies the path to the file that contains certificate revocation lists in PEM format.
	SSL_CRL = "ssl-crl"

	// SSL_CRLPATH specifies the path to the directory that contains certificate revocation-list files in PEM format.
	SSL_CRLPATH = "ssl-crlpath"

	// SSL_KEY specifies the path to the X.509 key file in PEM format.
	SSL_KEY = "ssl-crl"

	// TLS_VERSION specifies the TLS protocols permitted for classic MySQL protocol encrypted connections.
	TLS_VERSION = "tls-version"

	// AUTO_METHOD specifies the authentication method to use for the connection.
	// The default is AUTO
	AUTO_METHOD = "auth-method"

	// CONNECT_TIMEOUT defines an integer value used to configure the number of seconds that clients,
	// such as MySQL Shell, wait until they stop trying to connect to an unresponsive MySQL server.
	CONNECT_TIMEOUT = "connect_timeout"

	COMPRESSION = "compression"

	PROGRAM_NAME = "program_name"

	COLLATION = "collation"
)
