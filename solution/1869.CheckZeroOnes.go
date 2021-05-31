package solution

func CheckZeroOnes(s string) bool {
	max := [2]int{0, 0}
	tmp := [2]int{0, 0}
	p := uint8('0')
	for i := range s {
		if s[i] == p {
			tmp[p-'0']++
		} else {
			if tmp[p-'0'] > max[p-'0'] {
				max[p-'0'] = tmp[p-'0']
			}
			tmp[p-'0'] = 0
			p = s[i]
			tmp[p-'0']++
		}
	}
	for i := range tmp {
		if tmp[i] > max[i] {
			max[i] = tmp[i]
		}
	}
	return max[1] > max[0]
}
