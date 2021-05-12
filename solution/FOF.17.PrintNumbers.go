package solution

import "math"

func PrintNumbers(n int) []int {
	res := make([]int, int(math.Pow(10, float64(n))-1))
	for i := 1; i <= len(res); i++ {
		res[i-1] = i
	}
	return res
}
