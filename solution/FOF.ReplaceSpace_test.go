package solution

import "testing"

func TestReplaceSpace(t *testing.T) {
	s := "We are happy."
	exp := "We%20are%20happy."
	res := ReplaceSpace(s)
	if len(exp) != len(res) {
		t.Fail()
	}
	for i := range exp {
		if exp[i] != res[i] {
			t.Fail()
		}
	}
}
