package solution

import "testing"

func TestBusyStudent(t *testing.T) {
	s := []int{1, 2, 3}
	e := []int{3, 2, 7}
	queryTime := 4
	exp := 1
	if BusyStudent(s, e, queryTime) != exp {
		t.Fail()
	}
}
