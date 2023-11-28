package tests

import (
	"fmt"
	"testing"

	"github.com/leonelquinteros/gotext"
)

func TestGettext_1(t *testing.T) {
	// Set PO content
	str := `
msgid ""
msgstr ""

# Header below
"Plural-Forms: nplurals=2; plural=(n != 1);\n"

msgid "Translate this"
msgstr "Translated text"

msgid "Another string"
msgstr ""

msgid "One with var: %s"
msgid_plural "Several with vars: %s"
msgstr[0] "This one is the singular: %s"
msgstr[1] "This one is the plural: %s"
`

	// Create Po object
	po := gotext.NewPo()
	po.Parse([]byte(str))

	v1 := "Variable1"
	//v2 := "Variable2"

	fmt.Println(po.GetN("One with var: %s", "Several with vars: %s", 54, v1))
}

func TestGettext_Po_2(t *testing.T) {
	po := gotext.NewPo()
	po.ParseFile("./testdata/zh-Hans.po")

	str1 := po.Get("My text")
	fmt.Println(str1)

	str2 := po.Get("location.name")
	fmt.Println(str2)
}

func TestGettext_Mo_1(t *testing.T) {
	mo := gotext.NewMo()
	mo.ParseFile("./testdata/zh-Hans.mo")

	str1 := mo.Get("My text")
	fmt.Println(str1)

	str2 := mo.Get("location.name")
	fmt.Println(str2)
}
