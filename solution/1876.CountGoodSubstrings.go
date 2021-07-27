package solution

func CountGoodSubstrings(s string) int {
	res := 0
	for i := 0; i <= len(s)-3; i++ {
		if s[i] != s[i+1] && s[i] != s[i+2] && s[i+1] != s[i+2] {
			res++
		}
	}
	return res
}
