package solution

func MaxDepth(s string) int {
	res := 0
	tmp := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			tmp++
			if tmp > res {
				res = tmp
			}
		} else if s[i] == ')' {
			tmp--
		}
	}
	return res
}
