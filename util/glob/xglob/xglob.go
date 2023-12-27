package xglob

import (
	"github.com/gobwas/glob"
)

type Glob = glob.Glob

var (
	Compile = glob.Compile

	MustCompile = glob.MustCompile

	QuoteMeta = glob.QuoteMeta
)
