package solution

import "strings"

func MaxRepeating(sequence string, word string) int {
	for i := 0; i < len(sequence); i++ {
		if strings.Index(sequence, strings.Repeat(word, i+1)) == -1 {
			return i
		}
	}
	return len(sequence) / len(word)
}
