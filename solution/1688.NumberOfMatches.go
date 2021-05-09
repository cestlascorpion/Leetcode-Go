package solution

func NumberOfMatches(n int) int {
	sum := 0
	for n != 1 {
		if n%2 != 0 {
			sum += (n - 1) / 2
			n -= (n - 1) / 2
		} else {
			sum += n / 2
			n -= n / 2
		}
	}
	return sum
}
