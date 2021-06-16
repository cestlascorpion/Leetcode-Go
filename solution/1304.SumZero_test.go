package solution

import "testing"

func TestSumZero(t *testing.T) {
	n := 10
	res := SumZero(n)
	if len(res) != n {
		t.Fail()
	}
	sum := 0
	for i := range res {
		sum += res[i]
	}
	if sum != 0 {
		t.Fail()
	}
}
