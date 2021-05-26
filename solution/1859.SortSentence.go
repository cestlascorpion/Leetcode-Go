package solution

import (
	"strings"
)

func SortSentence(s string) string {
	slice := strings.Split(s, " ")
	dict := make([]string, len(slice))
	for i := range slice {
		idx := slice[i][len(slice[i])-1] - '1'
		dict[idx] = slice[i]
	}
	sentence := make([]uint8, 0)
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(dict[i])-1; j++ {
			sentence = append(sentence, dict[i][j])
		}
		if i < len(slice)-1 {
			sentence = append(sentence, ' ')
		}
	}
	return string(sentence)
}
