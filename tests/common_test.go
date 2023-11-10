package tests

import (
	"os"
	"path/filepath"
)

func GetTestdataPath() string {
	wd, _ := os.Getwd()
	return filepath.Join(wd, "testdata")
}
