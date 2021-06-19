package solution

import "testing"

func TestDaysBetweenDate(t *testing.T) {
	date1 := "2019-06-29"
	date2 := "2019-06-30"
	exp := 1
	if exp != DaysBetweenDates(date1, date2) {
		t.Fail()
	}
}

func TestDaysBetweenDate2(t *testing.T) {
	date1 := "2020-01-15"
	date2 := "2019-12-31"
	exp := 15
	if exp != DaysBetweenDates(date1, date2) {
		t.Fail()
	}
}
