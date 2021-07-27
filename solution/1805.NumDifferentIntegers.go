package solution

import "strings"

func isNumber(word *string, idx int) bool {
	return (*word)[idx] >= '0' && (*word)[idx] <= '9'
}

func NumDifferentIntegers(word string) int {
	set := make(map[string]int)
	idx := 0
	for idx < len(word) {
		for idx < len(word) {
			if isNumber(&word, idx) {
				break
			}
			idx++
		}
		if idx == len(word) {
			break
		}
		end := idx + 1
		for end < len(word) {
			if !isNumber(&word, end) {
				break
			}
			end++
		}
		num := word[idx:end]
		num = strings.TrimLeft(num, "0")
		set[num]++
		idx = end + 1
	}
	return len(set)
}
