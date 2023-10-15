package locales

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
	"testing"

	"github.com/karrick/godirwalk"
	"github.com/stretchr/testify/require"
)

func TestDir_1(t *testing.T) {
	err := fs.WalkDir(EmbedFS, ".", func(path string, d fs.DirEntry, err error) error {
		fmt.Printf("path:%s, info:%#v\n", path, d.Type().String())
		return nil
	})
	require.NoError(t, err)
}

func TestDir_2(t *testing.T) {
	err := filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		fmt.Printf("path:%s, info:%#v\n", path, d.Type().String())
		return nil
	})
	require.NoError(t, err)

	err1 := godirwalk.Walk(".", &godirwalk.Options{
		Callback: func(osPathname string, de *godirwalk.Dirent) error {
			// Following string operation is not most performant way
			// of doing this, but common enough to warrant a simple
			// example here:
			if strings.Contains(osPathname, ".git") {
				return godirwalk.SkipThis
			}
			fmt.Printf("%s %s\n", de.ModeType(), osPathname)
			return nil
		},
		Unsorted: true, // (optional) set true for faster yet non-deterministic enumeration (see godoc)
	})
	require.NoError(t, err1)
}
