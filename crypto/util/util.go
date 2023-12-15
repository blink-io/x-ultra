package util

import (
	"github.com/go-crypt/crypt"
)

type Digest = crypt.Digest

var (
	CheckPassword              = crypt.CheckPassword
	CheckPasswordWithPlainText = crypt.CheckPasswordWithPlainText
)
