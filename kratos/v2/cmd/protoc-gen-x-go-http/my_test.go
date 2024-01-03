package main

import (
	"fmt"
	"go/token"
	"strings"
	"testing"
	"unicode"
	"unicode/utf8"
)

func TestStr(t *testing.T) {
	pkgname := "github.com/go-kratos/kratos/v2/transport/http"
	// GoSanitized converts a string to a valid Go identifier.
	GoSanitized := func(s string) string {
		// Sanitize the input to the set of valid characters,
		// which must be '_' or be in the Unicode L or N categories.
		s = strings.Map(func(r rune) rune {
			if unicode.IsLetter(r) || unicode.IsDigit(r) {
				return r
			}
			return '_'
		}, s)

		// Prepend '_' in the event of a Go keyword conflict or if
		// the identifier is invalid (does not start in the Unicode L category).
		r, _ := utf8.DecodeRuneInString(s)
		if token.Lookup(s).IsKeyword() || !unicode.IsLetter(r) {
			return "_" + s
		}
		return s
	}

	// github_com_go_kratos_kratos_v2_transport_http
	nn := GoSanitized(pkgname)

	fmt.Println(nn)
}
