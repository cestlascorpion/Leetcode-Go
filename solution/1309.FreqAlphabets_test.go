package solution

import "testing"

func TestFreqAlphabets(t *testing.T) {
	s := "10#11#12"
	exp := "jkab"
	res := FreqAlphabets(s)
	if !IsSameString(exp, res) {
		t.Fail()
	}
}
