package solution

import "testing"

func TestAverage(t *testing.T) {
	salary := []int{4000, 3000, 1000, 2000}
	exp := 2500.0
	if Average(salary) != exp {
		t.Fail()
	}
}
