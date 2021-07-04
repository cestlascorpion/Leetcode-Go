package solution

func FinalPrices(prices []int) []int {
	res := make([]int, len(prices))
	for i := range prices {
		cut := 0
		for j := i + 1; j < len(prices); j++ {
			if prices[j] <= prices[i] {
				cut = prices[j]
				break
			}
		}
		res[i] = prices[i] - cut
	}
	return res
}
