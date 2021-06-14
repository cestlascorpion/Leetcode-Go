package solution

import (
	"testing"
)

func TestReplaceDigits(t *testing.T) {
	s := "a1c1e1"
	exp := "abcdef"
	res := ReplaceDigits(s)
	if !IsSameString(exp, res) {
		t.Fail()
	}
}
