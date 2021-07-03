package solution

func isPrefix(sentence, word *string, pos int) bool {
	for i := range *word {
		if i+pos >= len(*sentence) {
			return false
		}
		if (*sentence)[pos+i] != (*word)[i] {
			return false
		}
	}
	return true
}

func IsPrefixOfWord(sentence string, searchWord string) int {
	cnt := 0
	for i := range sentence {
		if i == 0 || sentence[i-1] == ' ' {
			cnt++
			if isPrefix(&sentence, &searchWord, i) {
				return cnt
			}
		}

	}
	return -1
}
