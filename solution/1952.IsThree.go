package solution

func IsThree(n int) bool {
	count := 0
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			if n/i == i {
				count++
			} else {
				count += 2
			}
		}
		if count > 3 {
			break
		}
	}
	return count == 3
}
