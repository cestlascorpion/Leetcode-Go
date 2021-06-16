package solution

func noZeroInteger(n int) bool {
	if n == 0 {
		return false
	}
	for n > 0 {
		if n%10 == 0 {
			return false
		}
		n = n / 10
	}
	return true
}

func GetNoZeroIntegers(n int) []int {
	res := make([]int, 2)
	for i := 1; i < n; i++ {
		if noZeroInteger(i) && noZeroInteger(n-i) {
			res[0] = i
			res[1] = n - i
			return res
		}
	}
	return res
}
