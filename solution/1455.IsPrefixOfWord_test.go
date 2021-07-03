package solution

import "testing"

func TestIsPrefixOfWord(t *testing.T) {
	sentence := "i love eating burger"
	searchWord := "burg"
	exp := 4
	if IsPrefixOfWord(sentence, searchWord) != exp {
		t.Fail()
	}
}


func TestIsPrefixOfWord2(t *testing.T) {
	sentence := "this problem is an easy problem"
	searchWord := "pro"
	exp := 2
	if IsPrefixOfWord(sentence, searchWord) != exp {
		t.Fail()
	}
}

func TestIsPrefixOfWord3(t *testing.T) {
	sentence := "i am tired"
	searchWord := "burg"
	exp := -1
	if IsPrefixOfWord(sentence, searchWord) != exp {
		t.Fail()
	}
}

func TestIsPrefixOfWord4(t *testing.T) {
	sentence := "hello from the other side"
	searchWord := "they"
	exp := -1
	if IsPrefixOfWord(sentence, searchWord) != exp {
		t.Fail()
	}
}
