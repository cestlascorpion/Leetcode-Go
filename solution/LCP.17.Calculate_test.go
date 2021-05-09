package solution

import "testing"

func TestCalculate(t *testing.T) {
	s := "AB"
	exp := 4
	if Calculate(s) != exp {
		t.Fail()
	}
}
