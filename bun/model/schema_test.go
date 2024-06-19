package model

import (
	"github.com/sanity-io/litter"
	"testing"
)

func TestSchemaColumn_1(t *testing.T) {
	type sch struct {
		Col Column
	}
	s := &sch{
		Col: "EL",
	}
	ss := litter.Sdump(s)
	println(ss)
}
