package solution

func HalvesAreAlike(s string) bool {
	n := len(s) / 2
	count := 0
	for i := 0; i < n; i++ {
		switch s[i] {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			count++
		}
		switch s[i+n] {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			count--
		}
	}
	return count == 0
}
