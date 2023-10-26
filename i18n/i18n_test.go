package i18n

import (
	"fmt"
	"testing"

	"golang.org/x/text/language"
)

func TestLangTag_1(t *testing.T) {
	var lt language.Tag
	if lt == language.Und {
		fmt.Println("Und: ------> ", lt.String())
	}

	fmt.Println(lt)
}
