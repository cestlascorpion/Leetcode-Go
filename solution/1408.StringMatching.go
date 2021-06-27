package solution

import "strings"

func StringMatching(words []string) []string {
	res := make([]string, 0)
	str := strings.Join(words, "/")
	for i := range words {
		if strings.Index(str, words[i]) != strings.LastIndex(str, words[i]) {
			res = append(res, words[i])
		}
	}
	return res
}
