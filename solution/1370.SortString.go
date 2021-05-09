package solution

func forward(dict []int, s []uint8, cursor int) int {
	for i := 0; i < len(dict); i++ {
		if dict[i] > 0 {
			s[cursor] = 'a' + (uint8(i))
			cursor++
			dict[i]--
		}
	}
	return cursor
}

func backward(dict []int, s []uint8, cursor int) int {
	for i := len(dict) - 1; i >= 0; i-- {
		if dict[i] > 0 {
			s[cursor] = 'a' + (uint8(i))
			cursor++
			dict[i]--
		}
	}
	return cursor
}

func SortString(s string) string {
	dict := make([]int, 26)
	for i := range s {
		dict[s[i]-'a']++
	}

	str := make([]uint8, len(s))
	cur := 0
	for cur != len(s) {
		for f := forward(dict, str, cur); f > cur; cur = f {
		}
		for b := backward(dict, str, cur); b > cur; cur = b {
		}
	}
	return string(str)
}
