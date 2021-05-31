package solution

func IsPowerOfFour(n int) bool {
	if n <= 0 {
		return false
	}
	base := 1
	for base < n {
		base = base << 2
	}
	if base == n {
		return true
	}
	return false
}
