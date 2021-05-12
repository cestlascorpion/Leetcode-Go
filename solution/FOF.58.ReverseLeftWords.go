package solution

import "strings"

func ReverseLeftWords(s string, n int) string {
	cur := n % len(s)
	if cur == 0 {
		return s
	}
	return strings.Join([]string{s[cur:], s[:cur]}, "")
}
