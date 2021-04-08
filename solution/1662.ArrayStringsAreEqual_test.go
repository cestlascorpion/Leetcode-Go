package solution

import "testing"

func TestArrayStringsAreEqual(t *testing.T) {
	word1 := []string{"ab", "c"}
	word2 := []string{"a", "bc"}
	if true != ArrayStringsAreEqual(word1, word2) {
		t.Fail()
	}
}
