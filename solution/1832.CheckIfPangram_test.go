package solution

import "testing"

func TestCheckIfPangram(t *testing.T) {
	sentence := "thequickbrownfoxjumpsoverthelazydog"
	exp := true
	if CheckIfPangram(sentence) != exp {
		t.Fail()
	}
}

func TestCheckIfPangram2(t *testing.T) {
	sentence := "leetcode"
	exp := false
	if CheckIfPangram(sentence) != exp {
		t.Fail()
	}
}
