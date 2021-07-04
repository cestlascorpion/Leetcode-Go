package solution

import "testing"

func TestFinalPrices(t *testing.T) {
	prices := []int{8, 4, 6, 2, 3}
	exp := []int{4, 2, 4, 2, 3}
	res := FinalPrices(prices)
	if !IsSameArray(exp, res) {
		t.Fail()
	}
}
