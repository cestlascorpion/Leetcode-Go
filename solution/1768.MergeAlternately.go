package solution

func MergeAlternately(word1 string, word2 string) string {
	l1, l2 := len(word1), len(word2)
	word := make([]byte, l1+l2)
	for i, j, k := 0, 0, 0; k < len(word); {
		if i < l1 {
			word[k] = word1[i]
			i++
			k++
		}
		if j < l2 {
			word[k] = word2[j]
			j++
			k++
		}
	}
	return string(word)
}
