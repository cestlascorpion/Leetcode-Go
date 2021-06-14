package solution

import "strings"

func FindOcurrences(text string, first string, second string) []string {
	res := make([]string, 0)
	words := strings.Split(text, " ")
	for i := range words {
		if i+2 < len(words) {
			if words[i] == first && words[i+1] == second {
				res = append(res, words[i+2])
			}
		}
	}
	return res
}
