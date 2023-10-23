package errcode

import (
	"fmt"
	"testing"

	"github.com/jackc/pgerrcode"
)

func TestErrcode_1(t *testing.T) {
	fmt.Println(pgerrcode.DeprecatedFeature)
}
