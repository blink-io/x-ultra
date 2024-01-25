package postgres

import (
	"errors"
	"fmt"
	"net"
	"net/url"
	"strings"
)

var asciiSpace = [256]uint8{'\t': 1, '\n': 1, '\v': 1, '\f': 1, '\r': 1, ' ': 1}

func isIPOnly(host string) bool {
	return net.ParseIP(strings.Trim(host, "[]")) != nil || !strings.Contains(host, ":")
}

type Settings map[string]string

func NewSettings() Settings {
	return make(Settings)
}

func (s Settings) Set(key, val string) {
	s[key] = val
}

func (s Settings) ToDSN() (string, error) {
	return "", nil
}

func (s Settings) ToURL() (string, error) {
	return "", nil
}

func (s Settings) ParseDSN(dsn string) error {
	if s == nil {
		return errors.New("settings is nil")
	}
	_, err := doParseDSN(dsn, s)
	return err
}

func (s Settings) ParseURL(urlstr string) error {
	if s == nil {
		return errors.New("settings is nil")
	}
	_, err := doParseURL(urlstr, s)
	return err
}

func ToDSN(s Settings) (string, error) {
	return s.ToDSN()
}

func ToURL(s Settings) (string, error) {
	return s.ToURL()
}

func ParseURL(urlstr string) (Settings, error) {
	return doParseURL(urlstr, nil)
}

func doParseURL(urlstr string, settings Settings) (Settings, error) {
	if settings == nil {
		settings = make(Settings)
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

func ParseDSN(dsn string) (Settings, error) {
	return doParseDSN(dsn, nil)
}

func doParseDSN(dsn string, settings Settings) (Settings, error) {
	if settings == nil {
		settings = make(Settings)
	}

	nameMap := map[string]string{
		"dbname": "database",
	}

	for len(dsn) > 0 {
		var key, val string
		eqIdx := strings.IndexRune(dsn, '=')
		if eqIdx < 0 {
			return nil, errors.New("invalid dsn")
		}

		key = strings.Trim(dsn[:eqIdx], " \t\n\r\v\f")
		dsn = strings.TrimLeft(dsn[eqIdx+1:], " \t\n\r\v\f")
		if len(dsn) == 0 {
		} else if dsn[0] != '\'' {
			end := 0
			for ; end < len(dsn); end++ {
				if asciiSpace[dsn[end]] == 1 {
					break
				}
				if dsn[end] == '\\' {
					end++
					if end == len(dsn) {
						return nil, errors.New("invalid backslash")
					}
				}
			}
			val = strings.Replace(strings.Replace(dsn[:end], "\\\\", "\\", -1), "\\'", "'", -1)
			if end == len(dsn) {
				dsn = ""
			} else {
				dsn = dsn[end+1:]
			}
		} else { // quoted string
			dsn = dsn[1:]
			end := 0
			for ; end < len(dsn); end++ {
				if dsn[end] == '\'' {
					break
				}
				if dsn[end] == '\\' {
					end++
				}
			}
			if end == len(dsn) {
				return nil, errors.New("unterminated quoted string in connection info string")
			}
			val = strings.Replace(strings.Replace(dsn[:end], "\\\\", "\\", -1), "\\'", "'", -1)
			if end == len(dsn) {
				dsn = ""
			} else {
				dsn = dsn[end+1:]
			}
		}

		if k, ok := nameMap[key]; ok {
			key = k
		}

		if key == "" {
			return nil, errors.New("invalid dsn")
		}

		settings[key] = val
	}

	return settings, nil
}

func defaultSettings() Settings {
	settings := NewSettings()
	return settings
}
