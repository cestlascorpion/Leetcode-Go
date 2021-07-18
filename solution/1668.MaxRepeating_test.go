package solution

import "testing"

func TestMaxRepeating(t *testing.T) {
	sequence := "ababc"
	word := "ab"
	exp := 2
	if MaxRepeating(sequence, word) != exp {
		t.Fail()
	}
}

func TestMaxRepeating2(t *testing.T) {
	sequence := "ababc"
	word := "ba"
	exp := 1
	if MaxRepeating(sequence, word) != exp {
		t.Fail()
	}
}

func TestMaxRepeating3(t *testing.T) {
	sequence := "ababc"
	word := "ac"
	exp := 0
	if MaxRepeating(sequence, word) != exp {
		t.Fail()
	}
}
