package solution

import "testing"

func TestReformatDate(t *testing.T) {
	date := "20th Oct 2052"
	exp := "2052-10-20"
	if ReformatDate(date) != exp {
		t.Fail()
	}
}
