package sqlite

import (
	"fmt"
	"testing"

	"github.com/stephenafamo/scan"
	"github.com/stephenafamo/scan/stdscan"
)

func TestSqlite_Scan_Select_1(t *testing.T) {
	db := getSqliteDB()
	apps, _ := stdscan.All(ctx, db, scan.StructMapper[*Application](), "SELECT * FROM applications")
	fmt.Println(apps)
}
