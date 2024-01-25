package postgres

import (
	"errors"
	"fmt"
	"net"
	"net/url"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

var asciiSpace = [256]uint8{'\t': 1, '\n': 1, '\v': 1, '\f': 1, '\r': 1, ' ': 1}

func isIPOnly(host string) bool {
	return net.ParseIP(strings.Trim(host, "[]")) != nil || !strings.Contains(host, ":")
}

type Settings map[string]string

func NewSettings() Settings {
	settings := defaultSettings()
	return settings
}

func (ss Settings) Set(key, val string) {
	ss[key] = val
}

func (ss Settings) ToDSN() (string, error) {
	return "", nil
}

func (ss Settings) ToURL() (string, error) {
	return "", nil
}

func (ss Settings) ParseDSN(dsnstr string) error {
	if ss == nil {
		return errors.New("settings is nil")
	}
	_, err := doParseDSN(dsnstr, ss)
	return err
}

func (ss Settings) ParseURL(urlstr string) error {
	if ss == nil {
		return errors.New("settings is nil")
	}
	_, err := doParseURL(urlstr, ss)
	return err
}

func ToDSN(s Settings) (string, error) {
	return s.ToDSN()
}

func ToURL(s Settings) (string, error) {
	return s.ToURL()
}

func ParseURL(urlstr string) (Settings, error) {
	return doParseURL(urlstr, defaultSettings())
}

func doParseURL(urlstr string, settings Settings) (Settings, error) {
	if settings == nil {
		settings = NewSettings()
	}

	url, err := url.Parse(urlstr)
	if err != nil {
		return nil, err
	}

	if url.User != nil {
		settings["user"] = url.User.Username()
		if password, present := url.User.Password(); present {
			settings["password"] = password
		}
	}

	// Handle multiple host:port's in url.Host by splitting them into host,host,host and port,port,port.
	var hosts []string
	var ports []string
	for _, host := range strings.Split(url.Host, ",") {
		if host == "" {
			continue
		}
		if isIPOnly(host) {
			hosts = append(hosts, strings.Trim(host, "[]"))
			continue
		}
		h, p, err := net.SplitHostPort(host)
		if err != nil {
			return nil, fmt.Errorf("failed to split host:port in '%s', err: %w", host, err)
		}
		if h != "" {
			hosts = append(hosts, h)
		}
		if p != "" {
			ports = append(ports, p)
		}
	}
	if len(hosts) > 0 {
		settings["host"] = strings.Join(hosts, ",")
	}
	if len(ports) > 0 {
		settings["port"] = strings.Join(ports, ",")
	}

	database := strings.TrimLeft(url.Path, "/")
	if database != "" {
		settings["database"] = database
	}

	nameMap := map[string]string{
		"dbname": "database",
	}

	for k, v := range url.Query() {
		if k2, present := nameMap[k]; present {
			k = k2
		}

		settings[k] = v[0]
	}

	return settings, nil
}

func ParseDSN(dsnstr string) (Settings, error) {
	return doParseDSN(dsnstr, defaultSettings())
}

func doParseDSN(dsnstr string, settings Settings) (Settings, error) {
	if settings == nil {
		settings = NewSettings()
	}

	nameMap := map[string]string{
		"dbname": "database",
	}

	for len(dsnstr) > 0 {
		var key, val string
		eqIdx := strings.IndexRune(dsnstr, '=')
		if eqIdx < 0 {
			return nil, errors.New("invalid DSN")
		}

		key = strings.Trim(dsnstr[:eqIdx], " \t\n\r\v\f")
		dsnstr = strings.TrimLeft(dsnstr[eqIdx+1:], " \t\n\r\v\f")
		if len(dsnstr) == 0 {
		} else if dsnstr[0] != '\'' {
			end := 0
			for ; end < len(dsnstr); end++ {
				if asciiSpace[dsnstr[end]] == 1 {
					break
				}
				if dsnstr[end] == '\\' {
					end++
					if end == len(dsnstr) {
						return nil, errors.New("invalid backslash")
					}
				}
			}
			val = strings.Replace(strings.Replace(dsnstr[:end], "\\\\", "\\", -1), "\\'", "'", -1)
			if end == len(dsnstr) {
				dsnstr = ""
			} else {
				dsnstr = dsnstr[end+1:]
			}
		} else { // quoted string
			dsnstr = dsnstr[1:]
			end := 0
			for ; end < len(dsnstr); end++ {
				if dsnstr[end] == '\'' {
					break
				}
				if dsnstr[end] == '\\' {
					end++
				}
			}
			if end == len(dsnstr) {
				return nil, errors.New("unterminated quoted string in connection info string")
			}
			val = strings.Replace(strings.Replace(dsnstr[:end], "\\\\", "\\", -1), "\\'", "'", -1)
			if end == len(dsnstr) {
				dsnstr = ""
			} else {
				dsnstr = dsnstr[end+1:]
			}
		}

		if k, ok := nameMap[key]; ok {
			key = k
		}

		if key == "" {
			return nil, errors.New("invalid DSN")
		}

		settings[key] = val
	}

	return settings, nil
}

func defaultSettings() Settings {
	settings := make(Settings)

	settings["host"] = defaultHost()
	settings["port"] = "5432"

	// Default to the OS user name. Purposely ignoring err getting user name from
	// OS. The client application will simply have to specify the user in that
	// case (which they typically will be doing anyway).
	user, err := user.Current()
	if err == nil {
		settings["user"] = user.Username
		settings["passfile"] = filepath.Join(user.HomeDir, ".pgpass")
		settings["servicefile"] = filepath.Join(user.HomeDir, ".pg_service.conf")
		sslcert := filepath.Join(user.HomeDir, ".postgresql", "postgresql.crt")
		sslkey := filepath.Join(user.HomeDir, ".postgresql", "postgresql.key")
		if _, err := os.Stat(sslcert); err == nil {
			if _, err := os.Stat(sslkey); err == nil {
				// Both the cert and key must be present to use them, or do not use either
				settings["sslcert"] = sslcert
				settings["sslkey"] = sslkey
			}
		}
		sslrootcert := filepath.Join(user.HomeDir, ".postgresql", "root.crt")
		if _, err := os.Stat(sslrootcert); err == nil {
			settings["sslrootcert"] = sslrootcert
		}
	}

	settings["target_session_attrs"] = "any"

	return settings
}

// defaultHost attempts to mimic libpq's default host. libpq uses the default unix socket location on *nix and localhost
// on Windows. The default socket location is compiled into libpq. Since pgx does not have access to that default it
// checks the existence of common locations.
func defaultHost() string {
	candidatePaths := []string{
		"/var/run/postgresql", // Debian
		"/private/tmp",        // OSX - homebrew
		"/tmp",                // standard PostgreSQL
	}

	for _, path := range candidatePaths {
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}

	return "localhost"
}

func mergeSettings(settingSets ...Settings) Settings {
	settings := make(Settings)

	for _, si := range settingSets {
		for k, v := range si {
			settings[k] = v
		}
	}

	return settings
}
