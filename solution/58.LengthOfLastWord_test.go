package solution

import "testing"

func LengthOfLastWord(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	i := n - 1
	for s[i] == ' ' {
		i--
	}
	j := i
	for j >= 0 && s[j] != ' ' {
		j--
	}
	return i - j
}

func TestLengthOfLastWord(t *testing.T) {
	s := "Hello World"
	if LengthOfLastWord(s) != 5 {
		t.Fatalf("case Hello World failed.")
	}
	t.Log("case Hello World succeeded.")
}