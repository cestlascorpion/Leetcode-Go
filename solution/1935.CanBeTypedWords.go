package solution

import (
	"strings"
)

func CanBeTypedWords(text string, brokenLetters string) int {
	cnt := 0
	broken := [26]int{0}
	for i := range brokenLetters {
		broken[brokenLetters[i]-'a']++
	}
	words := strings.Split(text, " ")
	for i := range words {
		for j := range words[i] {
			if broken[words[i][j]-'a'] != 0 {
				break
			}
			if j == len(words[i])-1 {
				cnt++
			}
		}
	}
	return cnt
}
