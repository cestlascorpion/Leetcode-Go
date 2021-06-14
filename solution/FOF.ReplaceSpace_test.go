package solution

import "testing"

func TestReplaceSpace(t *testing.T) {
	s := "We are happy."
	exp := "We%20are%20happy."
	res := ReplaceSpace(s)
	if !IsSameString(exp, res) {
		t.Fail()
	}
}
