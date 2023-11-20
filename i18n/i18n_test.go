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

func TestURL_1(t *testing.T) {
	u1 := "https://www.xxx.com/lang/zh-Hans.yaml"
	u2 := "https://www.xxx.com/langs.yaml?lang=zh-Hans"
	l1, f1 := ParsePath(u1)
	l2, f2 := ParsePath(u2)

	fmt.Println("lang: ", l1, ", format: ", f1)
	fmt.Println("lang: ", l2, ", format: ", f2)
}
