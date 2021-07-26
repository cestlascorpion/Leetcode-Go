package solution

func CheckOnesSegment(s string) bool {
	change := 0
	prev := byte('1')
	for i := range s {
		if s[i] != prev {
			change++
			prev = s[i]
		}
	}
	return change < 2
}
