package solution

func HammingDistance(x int, y int) int {
	res := 0
	for x > 0 || y > 0 {
		if (x % 2) != (y % 2) {
			res++
		}
		x = x / 2
		y = y / 2
	}
	return res
}
