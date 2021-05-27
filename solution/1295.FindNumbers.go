package solution

func findingsHelper(num int) int {
	res := 1
	for num > 9 {
		res++
		num = num / 10
	}
	return res
}

func FindNumbers(nums []int) int {
	res := 0
	for _, num := range nums {
		if (findingsHelper(num) & 1) == 0 {
			res++
		}
	}
	return res
}
