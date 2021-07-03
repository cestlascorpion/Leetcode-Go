package solution

func MaxPower(s string) int {
	power := 0
	pre := byte(0)
	cnt := 0
	for i := range s {
		if pre != s[i] {
			if cnt > power {
				power = cnt
			}
			pre = s[i]
			cnt = 1
		} else {
			cnt++
		}
	}
	if cnt > power {
		power = cnt
	}
	return power
}
