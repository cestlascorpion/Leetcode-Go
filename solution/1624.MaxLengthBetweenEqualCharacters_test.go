package solution

import "testing"

func TestMaxLengthBetweenEqualCharacters(t *testing.T) {
	s := "aa"
	exp := 0
	if MaxLengthBetweenEqualCharacters(s) != exp {
		t.Fail()
	}
}

func TestMaxLengthBetweenEqualCharacters2(t *testing.T) {
	s := "abca"
	exp := 2
	if MaxLengthBetweenEqualCharacters(s) != exp {
		t.Fail()
	}
}

func TestMaxLengthBetweenEqualCharacters3(t *testing.T) {
	s := "cbzxy"
	exp := -1
	if MaxLengthBetweenEqualCharacters(s) != exp {
		t.Fail()
	}
}

func TestMaxLengthBetweenEqualCharacters4(t *testing.T) {
	s := "cabbac"
	exp := 4
	if MaxLengthBetweenEqualCharacters(s) != exp {
		t.Fail()
	}
}