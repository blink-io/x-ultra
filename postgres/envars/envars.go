package envars

// Source: https://www.postgresql.org/docs/current/libpq-envars.html
const (
	// PGHOST behaves the same as the host connection parameter.
	PGHOST = "PGHOST"

	// PGHOSTADDR behaves the same as the hostaddr connection parameter. This can be set instead of or in addition to PGHOST to avoid DNS lookup overhead.
	PGHOSTADDR = "PGHOSTADDR"

	// PGPORT behaves the same as the port connection parameter.
	PGPORT = "PGPORT"

	//PGDATABASE behaves the same as the dbname connection parameter.
	PGDATABASE = "PGDATABASE"

	// PGUSER behaves the same as the user connection parameter.
	PGUSER = "PGUSER"

	// PGPASSWORD behaves the same as the password connection parameter. Use of this environment variable is not recommended for security reasons, as some operating systems allow non-root users to see process environment variables via ps; instead consider using a password file (see Section 34.16).
	PGPASSWORD = "PGPASSWORD"

	// PGPASSFILE behaves the same as the passfile connection parameter.
	PGPASSFILE = "PGPASSFILE"

	// PGREQUIREAUTH behaves the same as the require_auth connection parameter.
	PGREQUIREAUTH = "PGREQUIREAUTH"

	// PGCHANNELBINDING behaves the same as the channel_binding connection parameter.
	PGCHANNELBINDING = "PGCHANNELBINDING"

	// PGSERVICE behaves the same as the service connection parameter.
	PGSERVICE = "PGSERVICE"

	// PGSERVICEFILE specifies the name of the per-user connection service file (see Section 34.17). Defaults to ~/.pg_service.conf, or %APPDATA%\postgresql\.pg_service.conf on Microsoft Windows.
	PGSERVICEFILE = "PGSERVICEFILE"

	// PGOPTIONS behaves the same as the options connection parameter.
	PGOPTIONS = "PGOPTIONS"

	// PGAPPNAME behaves the same as the application_name connection parameter.
	PGAPPNAME = "PGAPPNAME"

	// PGSSLMODE behaves the same as the sslmode connection parameter.
	PGSSLMODE = "PGSSLMODE"

	//PGREQUIRESSL behaves the same as the requiressl connection parameter. This environment variable is deprecated in favor of the PGSSLMODE variable; setting both variables suppresses the effect of this one.
	PGREQUIRESSL = "PGREQUIRESSL"

	// PGSSLCOMPRESSION behaves the same as the sslcompression connection parameter.
	PGSSLCOMPRESSION = "PGSSLCOMPRESSION"

	// PGSSLCERT behaves the same as the sslcert connection parameter.
	PGSSLCERT = "PGSSLCERT"

	// PGSSLKEY behaves the same as the sslkey connection parameter.
	PGSSLKEY = "PGSSLKEY"

	// PGSSLCERTMODE behaves the same as the sslcertmode connection parameter.
	PGSSLCERTMODE = "PGSSLCERTMODE"

	// PGSSLROOTCERT behaves the same as the sslrootcert connection parameter.
	PGSSLROOTCERT = "PGSSLROOTCERT"

	// PGSSLCRL behaves the same as the sslcrl connection parameter.
	PGSSLCRL = "PGSSLCRL"

	// PGSSLCRLDIR behaves the same as the sslcrldir connection parameter.
	PGSSLCRLDIR = "PGSSLCRLDIR"

	// PGSSLSNI behaves the same as the sslsni connection parameter.
	PGSSLSNI = "PGSSLSNI"

	// PGREQUIREPEER behaves the same as the requirepeer connection parameter.
	PGREQUIREPEER = "PGREQUIREPEER"

	// PGSSLMINPROTOCOLVERSION behaves the same as the ssl_min_protocol_version connection parameter.
	PGSSLMINPROTOCOLVERSION = "PGSSLMINPROTOCOLVERSION"

	// PGSSLMAXPROTOCOLVERSION behaves the same as the ssl_max_protocol_version connection parameter.
	PGSSLMAXPROTOCOLVERSION = "PGSSLMAXPROTOCOLVERSION"

	// PGGSSENCMODE behaves the same as the gssencmode connection parameter.
	PGGSSENCMODE = "PGGSSENCMODE"

	// PGKRBSRVNAME behaves the same as the krbsrvname connection parameter.
	PGKRBSRVNAME = "PGKRBSRVNAME"

	// PGGSSLIB behaves the same as the gsslib connection parameter.
	PGGSSLIB = "PGGSSLIB"

	// PGGSSDELEGATION behaves the same as the gssdelegation connection parameter.
	PGGSSDELEGATION = "PGGSSDELEGATION"

	// PGCONNECT_TIMEOUT behaves the same as the connect_timeout connection parameter.
	PGCONNECT_TIMEOUT = "PGCONNECT_TIMEOUT"

	// PGCLIENTENCODING behaves the same as the client_encoding connection parameter.
	PGCLIENTENCODING = "PGCLIENTENCODING"

	// PGTARGETSESSIONATTRS behaves the same as the target_session_attrs connection parameter.
	PGTARGETSESSIONATTRS = "PGTARGETSESSIONATTRS"

	// PGLOADBALANCEHOSTS behaves the same as the load_balance_hosts connection parameter.
	// The following environment variables can be used to specify default behavior for each PostgreSQL session.
	// (See also the ALTER ROLE and ALTER DATABASE commands for ways to set default behavior on a per-user or per-database basis.)
	PGLOADBALANCEHOSTS = "PGLOADBALANCEHOSTS"

	// PGDATESTYLE sets the default style of date/time representation. (Equivalent to SET datestyle TO ....)
	PGDATESTYLE = "PGDATESTYLE"

	// PGTZ sets the default time zone. (Equivalent to SET timezone TO ....)
	PGTZ = "PGTZ"

	// PGGEQO sets the default mode for the genetic query optimizer. (Equivalent to SET geqo TO ....)
	// Refer to the SQL command SET for information on correct values for these environment variables.
	// The following environment variables determine internal behavior of libpq; they override compiled-in defaults.
	PGGEQO = "PGGEQO"

	// PGSYSCONFDIR sets the directory containing the pg_service.conf file and in a future version possibly other system-wide configuration files.
	PGSYSCONFDIR = "PGSYSCONFDIR"

	// PGLOCALEDIR sets the directory containing the locale files for message localization.
	PGLOCALEDIR = "PGLOCALEDIR"
)
