package solution

func CountOdds(low int, high int) int {
	return ((high + 1) >> 1) - (low >> 1)
}
