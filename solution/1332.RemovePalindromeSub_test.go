package solution

import "testing"

func TestRemovePalindromeSub(t *testing.T) {
	s := []string{"ababa", "abb", ""}
	exp := []int{1, 2, 0}
	for i := range s {
		if exp[i] != RemovePalindromeSub(s[i]) {
			t.Fail()
		}
	}
}
