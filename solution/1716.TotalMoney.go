package solution

func TotalMoney(n int) int {
	base := 1
	sum := 0
	for n != 0 {
		for i := 0; i < 7; i++ {
			if n == 0 {
				break
			}
			sum += base + i
			n--
		}
		base++
	}
	return sum
}