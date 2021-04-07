package solution

func CountConsistentStrings(allowed string, words []string) int {
	res := len(words)
	var bucket [26]int
	for _, c := range allowed {
		bucket[c - 'a']++
	}
	for i := 0; i < len(words); i++ {
		var j int
		for ; j < len(words[i]); j++ {
			if bucket[words[i][j] - 'a'] == 0 {
				res --
				break
			}
		}
	}
	return res
}
