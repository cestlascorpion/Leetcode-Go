package solution

func ZeroOneMinOperations(s string) int {
	cnt := [2]int{0, 0} // 0, 1
	ch := [2]byte{'0', '1'}
	for i := range s {
		if s[i] == ch[i%2] {
			cnt[1]++
		} else {
			cnt[0]++
		}
	}
	if cnt[0] > cnt[1] {
		return cnt[1]
	}
	return cnt[0]
}
