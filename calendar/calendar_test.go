package calendar

import (
	"testing"
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/us"
)

func TestCal_1(t *testing.T) {
	cnCal := cal.NewBusinessCalendar()
	cnCal.AddHoliday(&cal.Holiday{})
	cnCal.AddHoliday(us.ChristmasDay)

	cnCal.IsHoliday(time.Now())
}
