package solution

func ballSum(num int) int {
	sum := 0
	for num > 0 {
		sum += num % 10
		num = num / 10
	}
	return sum
}

func CountBalls(lowLimit int, highLimit int) int {
	count := make(map[int]int)
	for i := lowLimit; i <= highLimit; i++ {
		count[ballSum(i)]++
	}
	res := 0
	for _, v := range count {
		if v > res {
			res = v
		}
	}
	return res
}
