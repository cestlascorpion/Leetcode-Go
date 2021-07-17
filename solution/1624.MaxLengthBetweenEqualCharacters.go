package solution

func MaxLengthBetweenEqualCharacters(s string) int {
	res := -1
	dict := [26]int{}
	for i := range dict {
		dict[i] = -1
	}
	for i := range s {
		if dict[s[i]-'a'] == -1 {
			dict[s[i]-'a'] = i
		} else {
			delta := i - dict[s[i]-'a'] - 1
			if delta > res {
				res = delta
			}
		}
	}
	return res
}
