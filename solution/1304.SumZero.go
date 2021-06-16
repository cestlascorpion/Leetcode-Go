package solution

func SumZero(n int) []int {
	res := make([]int, 0, n)
	if n%2 != 0 {
		res = append(res, 0)
	}
	for i := 1; i <= n/2; i++ {
		res = append(res, i, -i)
	}
	return res
}
