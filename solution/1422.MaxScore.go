package solution

func MaxScore(s string) int {
	sum := 0
	cnt := make([][2]int, len(s)+1)
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			cnt[i+1][0] = cnt[i][0] + 1
		} else {
			cnt[i+1][0] = cnt[i][0]
		}
		if s[len(s)-1-i] == '1' {
			cnt[len(cnt)-1-i-1][1] = cnt[len(cnt)-1-i][1] + 1
		} else {
			cnt[len(cnt)-1-i-1][1] = cnt[len(cnt)-1-i][1]
		}
	}
	for i := 1; i < len(cnt)-1; i++ {
		if s := cnt[i][0] + cnt[i][1]; sum < s {
			sum = s
		}
	}
	return sum
}
