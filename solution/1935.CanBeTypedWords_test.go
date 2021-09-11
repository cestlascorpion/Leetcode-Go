package solution

import "testing"

func TestCanBeTypedWords(t *testing.T) {
	text := "hello world"
	brokenLetters := "ad"
	if CanBeTypedWords(text, brokenLetters) != 1 {
		t.Failed()
	}
}

func TestCanBeTypedWords2(t *testing.T) {
	text := "leet code"
	brokenLetters := "lt"
	if CanBeTypedWords(text, brokenLetters) != 1 {
		t.Failed()
	}
}

func TestCanBeTypedWords3(t *testing.T) {
	text := "leet code"
	brokenLetters := "e"
	if CanBeTypedWords(text, brokenLetters) != 0 {
		t.Failed()
	}
}
