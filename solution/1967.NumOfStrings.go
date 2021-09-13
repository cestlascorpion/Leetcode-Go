package solution

import "strings"

func NumOfStrings(patterns []string, word string) int {
	ret := 0
	for i := range patterns {
		if strings.Contains(word, patterns[i]) {
			ret++
		}
	}
	return ret
}
