package pg

import (
	"os"
	"path/filepath"
	"time"

	"github.com/blink-io/x/bun"

	xsql "github.com/blink-io/x/sql"
)

func dbOpts() []bun.Option {
	opts := []bun.Option{}
	return opts
}

func pgCfg() *xsql.Config {
	cfg := &xsql.Config{
		Context:       ctx,
		Dialect:       xsql.DialectPostgres,
		Name:          "test",
		User:          "postgres",
		Port:          5432,
		Host:          "localhost",
		Loc:           time.Local,
		ValidationSQL: "SELECT 1;",
		ClientName:    "blink-x-pgx",
		Password:      "postgres",
		DriverHooks:   newDriverHooks(),
		Additions:     map[string]string{
			//xsql.AdditionUsePool: "true",
		},
	}
	return cfg
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
