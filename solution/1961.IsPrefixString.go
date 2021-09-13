package solution

func IsPrefixString(s string, words []string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(words) {
		for k := range words[j] {
			if i >= len(s) {
				return false
			}
			if s[i] != words[j][k] {
				return false
			}
			i++
		}
		j++
	}
	return i == len(s)
}
