package dotenv

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestDotenv_1(t *testing.T) {
	path1 := ""
	p1 := filepath.Join(path1, ".env")
	fmt.Print("p1:  ", p1)

	path2 := "hello"
	p2 := filepath.Join(path2, ".env")
	fmt.Print("p2:  ", p2)
}
