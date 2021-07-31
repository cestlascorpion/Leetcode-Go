package solution

func MakeEqual(words []string) bool {
	sum := make(map[byte]int)
	for i := range words {
		for j := range words[i] {
			sum[words[i][j]]++
		}
	}
	for _, v := range sum {
		if v%len(words) != 0 {
			return false
		}
	}
	return true
}
