package solution

import (
	"math"
	"time"
)

func DaysBetweenDates(date1 string, date2 string) int {
	layout:="2006-01-02"
	d1,_:=time.Parse(layout,date1)
	d2,_:=time.Parse(layout,date2)
	return int(math.Abs(d1.Sub(d2).Hours()/24))
}
