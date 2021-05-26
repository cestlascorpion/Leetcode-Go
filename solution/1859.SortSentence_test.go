package solution

import "testing"

func TestSortSentence(t *testing.T) {
	s := "is2 sentence4 This1 a3"
	exp := "This is a sentence"
	res := SortSentence(s)
	if len(exp) != len(res) {
		t.Fail()
	}
	for i := range exp {
		if exp[i] != res[i] {
			t.Fail()
		}
	}
}
