package solution

import "testing"

func TestIsSumEqual(t *testing.T) {
	firstWord := "acb"
	secondWord := "cba"
	targetWord := "cdb"
	exp := true
	if IsSumEqual(firstWord, secondWord, targetWord) != exp {
		t.Fail()
	}
}

func TestIsSumEqual2(t *testing.T) {
	firstWord := "aaa"
	secondWord := "a"
	targetWord := "aaaa"
	exp := true
	if IsSumEqual(firstWord, secondWord, targetWord) != exp {
		t.Fail()
	}
}

func TestIsSumEqual3(t *testing.T) {
	firstWord := "j"
	secondWord := "b"
	targetWord := "ba"
	exp := true
	if IsSumEqual(firstWord, secondWord, targetWord) != exp {
		t.Fail()
	}
}

func TestIsSumEqual4(t *testing.T) {
	firstWord := "h"
	secondWord := "fhjfdghj"
	targetWord := "fhjfdgig"
	exp := true
	if IsSumEqual(firstWord, secondWord, targetWord) != exp {
		t.Fail()
	}
}

