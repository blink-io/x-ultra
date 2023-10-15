package i18n

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/blink-io/x/i18n"
	"github.com/stretchr/testify/require"
)

func Test_Loader_1(t *testing.T) {
	bundle := i18n.Default()
	dir := filepath.Join(".", "locales")
	langs := []string{"zh-Hans", "en-US"}
	require.NoError(t, i18n.NewDirLoader(dir, i18n.DefaultSuffixes...).Load(bundle))

	for _, lang := range langs {
		tr := i18n.GetT(lang)

		msg := tr("hello", i18n.MD{"Name": "兜兜"}.O())
		fmt.Printf("msg: %s\n", msg)
	}
}
