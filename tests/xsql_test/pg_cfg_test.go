package xsql_test

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	xsql "github.com/blink-io/x/sql"
)

var pgOpt = &xsql.Options{
	Context:       ctx,
	Dialect:       xsql.DialectPostgres,
	Name:          "blink",
	User:          "blinkbot",
	Port:          15432,
	Host:          "192.168.11.179",
	Loc:           time.Local,
	ValidationSQL: "SELECT 1;",
	ClientName:    "blink-dev",
	Password:      getPgPwd(),
	DriverHooks:   newDriverHooks(),
	Logger: func(format string, args ...any) {
		msg := fmt.Sprintf(format, args...)
		slog.Default().Info(msg)
	},
}

func getPgPwd() string {
	homedir, _ := os.UserHomeDir()

	data, err := os.ReadFile(filepath.Join(homedir, ".passwd.pg"))
	if err != nil {
		panic(err)
	}

	// Remove \n
	pwd := string(data[:len(data)-1])
	return pwd
}

func getPgFuncsMap() map[string]string {
	funcsMap := map[string]string{
		"version":                  "version()",
		"gen_random_uuid":          "gen_random_uuid()",
		"current_database":         "current_database()",
		"inet_client_addr":         "inet_client_addr()",
		"inet_client_port":         "inet_client_port()",
		"inet_server_addr":         "inet_server_addr()",
		"inet_server_port":         "inet_server_port()",
		"pg_backend_pid":           "pg_backend_pid()",
		"session_user":             "session_user",
		"current_user":             "current_user",
		"pg_conf_load_time":        "pg_conf_load_time()",
		"pg_postmaster_start_time": "pg_postmaster_start_time()",
	}
	return funcsMap
}
