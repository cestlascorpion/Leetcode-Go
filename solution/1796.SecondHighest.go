package solution

func SecondHighest(s string) int {
	num := [10]int{}
	for i := range s {
		if s[i] >= '0' && s[i] <= '9' {
			num[s[i]-'0']++
		}
	}
	first := false
	for i := 9; i >= 0; i-- {
		if num[i] > 0 {
			if first {
				return i
			} else {
				first = true
			}
		}
	}
	return -1
}
