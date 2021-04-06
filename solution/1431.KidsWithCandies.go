package solution

func KidsWithCandies(candies []int, extraCandies int) []bool {
	res := make([]bool, len(candies))
	if len(candies) == 0 {
		return res
	}
	max := candies[0]
	for i := 1; i < len(candies); i++ {
		if candies[i] > max {
			max = candies[i]
		}
	}
	for i, v := range candies {
		if v+extraCandies >= max {
			res[i] = true
		} else {
			res[i] = false
		}
	}
	return res
}
