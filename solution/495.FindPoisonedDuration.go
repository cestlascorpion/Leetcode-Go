package solution

func FindPoisonedDuration(series []int, duration int) int {
	res := 0
	n := len(series)
	if duration == 0 || n == 0 {
		return res
	}
	for i := range series {
		if i+1 < n {
			if series[i]+duration < series[i+1] {
				res += duration
			} else {
				res += series[i+1] - series[i]
			}
		} else {
			res += duration
		}
	}
	return res
}
