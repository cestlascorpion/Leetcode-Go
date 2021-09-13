package solution

import "testing"

func TestGetLucky(t *testing.T) {
	s := "leetcode"
	k := 2
	exp := 6
	if GetLucky(s, k) != exp {
		t.Fail()
	}
}
