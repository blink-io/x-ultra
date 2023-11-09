package strfmt

import (
	"fmt"
	"testing"
	"time"
)

func TestDate_1(t *testing.T) {
	now := time.Now()

	tm := Date(now)
	dt := DateTime(now)
	fmt.Println("DateTime: ", dt)
	fmt.Println("Date    : ", tm)
}
