package solution

import "testing"

func TestReversePrefix(t *testing.T) {
	word := "abcdefd"
	exp := "dcbaefd"
	if ReversePrefix(word,'d') != exp {
		t.FailNow()
	}
}
