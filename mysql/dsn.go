package mysql

import (
	"context"
	"database/sql"
	"errors"
	"net"
	"strings"
	"time"

	"github.com/blink-io/x/cast"
)

const (
	defaultCollation        = "utf8mb4_general_ci"
	binaryCollation         = "binary"
	defaultMaxAllowedPacket = 64 << 20 // 64 MiB. See https://github.com/go-sql-driver/mysql/issues/1355
)

var (
	errInvalidDSNUnescaped       = errors.New("invalid DSN: did you forget to escape a param value?")
	errInvalidDSNAddr            = errors.New("invalid DSN: network address not terminated (missing closing brace)")
	errInvalidDSNNoSlash         = errors.New("invalid DSN: missing the slash separating the database name")
	errInvalidDSNUnsafeCollation = errors.New("invalid DSN: interpolateParams can not be used with unsafe collations")
)

func QueryVersion(ctx context.Context, queryRowContext func(ctx context.Context, query string, args ...any) *sql.Row) string {
	row := queryRowContext(ctx, "SELECT version() as version;")
	var str string
	_ = row.Scan(&str)
	return str
}

type Settings map[string]string

func (s Settings) Set(key, val string) {
	s[key] = val
}

func defaultSettings() Settings {
	settings := make(Settings)
	settings["Collation"] = defaultCollation
	settings["Loc"] = time.Local.String()
	settings["maxAllowedPacket"] = cast.ToString(defaultMaxAllowedPacket)
	settings["allowNativePasswords"] = "true"
	settings["checkConnLiveness"] = "true"
	return settings
}

// ParseDSN parses the DSN string to a Config
func ParseDSN(dsn string) (Settings, err error) {
	// New config with some default values
	settings := defaultSettings()

	// [user[:password]@][net[(addr)]]/dbname[?param1=value1&paramN=valueN]
	// Find the last '/' (since the password or the net addr might contain a '/')
	foundSlash := false
	for i := len(dsn) - 1; i >= 0; i-- {
		if dsn[i] == '/' {
			foundSlash = true
			var j, k int

			// left part is empty if i <= 0
			if i > 0 {
				// [username[:password]@][protocol[(address)]]
				// Find the last '@' in dsn[:i]
				for j = i; j >= 0; j-- {
					if dsn[j] == '@' {
						// username[:password]
						// Find the first ':' in dsn[:j]
						for k = 0; k < j; k++ {
							if dsn[k] == ':' {
								settings.Set("passwd", dsn[k+1:j])
								break
							}
						}
						settings.Set("user", dsn[:k])

						break
					}
				}

				// [protocol[(address)]]
				// Find the first '(' in dsn[j+1:i]
				for k = j + 1; k < i; k++ {
					if dsn[k] == '(' {
						// dsn[i-1] must be == ')' if an address is specified
						if dsn[i-1] != ')' {
							if strings.ContainsRune(dsn[k+1:i], ')') {
								return nil, errInvalidDSNUnescaped
							}
							return nil, errInvalidDSNAddr
						}
						settings.Set("addr", dsn[k+1:i-1])
						break
					}
				}
				settings.Set("net", dsn[j+1:k])
			}

			// dbname[?param1=value1&...&paramN=valueN]
			// Find the first '?' in dsn[i+1:]
			for j = i + 1; j < len(dsn); j++ {
				if dsn[j] == '?' {
					if err = parseDSNParams(settings, dsn[j+1:]); err != nil {
						return
					}
					break
				}
			}
			settings.Set("dbname", dsn[i+1:j])
			break
		}
	}

	if !foundSlash && len(dsn) > 0 {
		return nil, errInvalidDSNNoSlash
	}

	//if err = cfg.normalize(); err != nil {
	//	return nil, err
	//}
	return
}

// parseDSNParams parses the DSN "query string"
// Values must be url.QueryEscape'ed
func parseDSNParams(s Settings, params string) (err error) {
	//for _, v := range strings.Split(params, "&") {
	//	param := strings.SplitN(v, "=", 2)
	//	if len(param) != 2 {
	//		continue
	//	}
	//
	//	// cfg params
	//	switch value := param[1]; param[0] {
	//	// Disable INFILE allowlist / enable all files
	//	case "allowAllFiles":
	//		var isBool bool
	//		cfg.AllowAllFiles, isBool = readBool(value)
	//		if !isBool {
	//			return errors.New("invalid bool value: " + value)
	//		}
	//
	//	// Use cleartext authentication mode (MySQL 5.5.10+)
	//	case "allowCleartextPasswords":
	//		var isBool bool
	//		cfg.AllowCleartextPasswords, isBool = readBool(value)
	//		if !isBool {
	//			return errors.New("invalid bool value: " + value)
	//		}
	//
	//	// Allow fallback to unencrypted connection if server does not support TLS
	//	case "allowFallbackToPlaintext":
	//		var isBool bool
	//		cfg.AllowFallbackToPlaintext, isBool = readBool(value)
	//		if !isBool {
	//			return errors.New("invalid bool value: " + value)
	//		}
	//
	//	// Use native password authentication
	//	case "allowNativePasswords":
	//		var isBool bool
	//		cfg.AllowNativePasswords, isBool = readBool(value)
	//		if !isBool {
	//			return errors.New("invalid bool value: " + value)
	//		}
	//
	//	// Use old authentication mode (pre MySQL 4.1)
	//	case "allowOldPasswords":
	//		var isBool bool
	//		cfg.AllowOldPasswords, isBool = readBool(value)
	//		if !isBool {
	//			return errors.New("invalid bool value: " + value)
	//		}
	//
	//	// Check connections for Liveness before using them
	//	case "checkConnLiveness":
	//		var isBool bool
	//		cfg.CheckConnLiveness, isBool = readBool(value)
	//		if !isBool {
	//			return errors.New("invalid bool value: " + value)
	//		}
	//
	//	// Switch "rowsAffected" mode
	//	case "clientFoundRows":
	//		var isBool bool
	//		cfg.ClientFoundRows, isBool = readBool(value)
	//		if !isBool {
	//			return errors.New("invalid bool value: " + value)
	//		}
	//
	//	// Collation
	//	case "collation":
	//		cfg.Collation = value
	//
	//	case "columnsWithAlias":
	//		var isBool bool
	//		cfg.ColumnsWithAlias, isBool = readBool(value)
	//		if !isBool {
	//			return errors.New("invalid bool value: " + value)
	//		}
	//
	//	// Compression
	//	case "compress":
	//		return errors.New("compression not implemented yet")
	//
	//	// Enable client side placeholder substitution
	//	case "interpolateParams":
	//		var isBool bool
	//		cfg.InterpolateParams, isBool = readBool(value)
	//		if !isBool {
	//			return errors.New("invalid bool value: " + value)
	//		}
	//
	//	// Time Location
	//	case "loc":
	//		if value, err = url.QueryUnescape(value); err != nil {
	//			return
	//		}
	//		cfg.Loc, err = time.LoadLocation(value)
	//		if err != nil {
	//			return
	//		}
	//
	//	// multiple statements in one query
	//	case "multiStatements":
	//		var isBool bool
	//		cfg.MultiStatements, isBool = readBool(value)
	//		if !isBool {
	//			return errors.New("invalid bool value: " + value)
	//		}
	//
	//	// time.Time parsing
	//	case "parseTime":
	//		var isBool bool
	//		cfg.ParseTime, isBool = readBool(value)
	//		if !isBool {
	//			return errors.New("invalid bool value: " + value)
	//		}
	//
	//	// I/O read Timeout
	//	case "readTimeout":
	//		cfg.ReadTimeout, err = time.ParseDuration(value)
	//		if err != nil {
	//			return
	//		}
	//
	//	// Reject read-only connections
	//	case "rejectReadOnly":
	//		var isBool bool
	//		cfg.RejectReadOnly, isBool = readBool(value)
	//		if !isBool {
	//			return errors.New("invalid bool value: " + value)
	//		}
	//
	//	// Server public key
	//	case "serverPubKey":
	//		name, err := url.QueryUnescape(value)
	//		if err != nil {
	//			return fmt.Errorf("invalid value for server pub key name: %v", err)
	//		}
	//		cfg.ServerPubKey = name
	//
	//	// Strict mode
	//	case "strict":
	//		panic("strict mode has been removed. See https://github.com/go-sql-driver/mysql/wiki/strict-mode")
	//
	//	// Dial Timeout
	//	case "timeout":
	//		cfg.Timeout, err = time.ParseDuration(value)
	//		if err != nil {
	//			return
	//		}
	//
	//	// TLS-Encryption
	//	case "tls":
	//		boolValue, isBool := readBool(value)
	//		if isBool {
	//			if boolValue {
	//				cfg.TLSConfig = "true"
	//			} else {
	//				cfg.TLSConfig = "false"
	//			}
	//		} else if vl := strings.ToLower(value); vl == "skip-verify" || vl == "preferred" {
	//			cfg.TLSConfig = vl
	//		} else {
	//			name, err := url.QueryUnescape(value)
	//			if err != nil {
	//				return fmt.Errorf("invalid value for TLS config name: %v", err)
	//			}
	//			cfg.TLSConfig = name
	//		}
	//
	//	// I/O write Timeout
	//	case "writeTimeout":
	//		cfg.WriteTimeout, err = time.ParseDuration(value)
	//		if err != nil {
	//			return
	//		}
	//	case "maxAllowedPacket":
	//		cfg.MaxAllowedPacket, err = strconv.Atoi(value)
	//		if err != nil {
	//			return
	//		}
	//	default:
	//		// lazy init
	//		if cfg.Params == nil {
	//			cfg.Params = make(map[string]string)
	//		}
	//
	//		if cfg.Params[param[0]], err = url.QueryUnescape(value); err != nil {
	//			return
	//		}
	//	}
	//}

	return
}

func ensureHavePort(addr string) string {
	if _, _, err := net.SplitHostPort(addr); err != nil {
		return net.JoinHostPort(addr, "3306")
	}
	return addr
}
