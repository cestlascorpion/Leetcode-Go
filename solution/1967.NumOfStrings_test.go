package solution

import (
	"testing"
)

func TestNumOfStrings(t *testing.T) {
	patterns := []string{"a", "ab", "c"}
	word := "abdef"
	exp := 2
	if NumOfStrings(patterns, word) != exp {
		t.Failed()
	}
}
