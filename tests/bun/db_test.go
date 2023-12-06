package bun

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/blink-io/x/bun/extra/logging"
	"github.com/blink-io/x/bun/extra/timing"
	xsql "github.com/blink-io/x/sql"
	"github.com/stretchr/testify/require"
)

func GetPgDB() *xsql.DB {
	homedir, _ := os.UserHomeDir()

	data, err := os.ReadFile(filepath.Join(homedir, ".passwd.pg"))
	if err != nil {
		panic(err)
	}

	pwd := strings.TrimSuffix(string(data), "\n")

	opt := &xsql.Options{
		Context:       context.Background(),
		Dialect:       xsql.DialectPostgres,
		Name:          "blink",
		User:          "blinkbot",
		Port:          15432,
		Host:          "192.168.11.179",
		Password:      pwd,
		Loc:           time.Local,
		ValidationSQL: "SELECT 1;",
		ClientName:    "blink-dev",
	}

	db, err1 := xsql.NewDB(opt)
	if err1 != nil {
		panic(err1)
	}

	db.AddQueryHook(logging.Func(log.Printf))
	db.AddQueryHook(timing.New())

	return db
}

func TestPG_Connect_1(t *testing.T) {
	db := GetPgDB()
	row := db.QueryRow("select version();")

	var vinfo string
	err := row.Scan(&vinfo)
	require.NoError(t, err)

	fmt.Println("DB Version: ", vinfo)
}
