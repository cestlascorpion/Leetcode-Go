package solution

func MinCount(coins []int) int {
	res := 0
	for i := range coins {
		if (coins[i] & 1) == 1 {
			res += coins[i]/2 + 1
		} else {
			res += coins[i] / 2
		}
	}
	return res
}
