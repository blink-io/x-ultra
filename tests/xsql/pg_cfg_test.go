package bun

import (
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
		"version":          "version()",
		"gen_random_uuid":  "gen_random_uuid()",
		"current_database": "current_database()",
	}
	return funcsMap
}
