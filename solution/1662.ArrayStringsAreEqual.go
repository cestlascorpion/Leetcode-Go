package solution

func ArrayStringsAreEqual(word1 []string, word2 []string) bool {
	// s1 := strings.Join(word1, "")
	// s2 := strings.Join(word2, "")
	// return s1 == s2
	l1, l2 := len(word1), len(word2)
	index1, index2, wordIndex1, wordIndex2 := 0, 0, 0, 0
	for index1 < l1 && index2 < l2 {
		if word1[index1][wordIndex1] != word2[index2][wordIndex2] {
			return false
		}
		wordIndex1++
		wordIndex2++
		if wordIndex1 == len(word1[index1]) {
			index1++
			wordIndex1 = 0
		}
		if wordIndex2 == len(word2[index2]) {
			index2++
			wordIndex2 = 0
		}
	}
	return index1 == l1 && index2 == l2
}
