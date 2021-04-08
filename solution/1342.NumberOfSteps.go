package solution

func NumberOfSteps(num int) int {
	res := 0
	for num != 0 {
		if (num & 1) == 0 {
			num /= 2
		} else {
			num--
		}
		res++
	}
	return res
}
