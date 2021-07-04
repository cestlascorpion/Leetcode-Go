package solution

func Average(salary []int) float64 {
	num := len(salary) - 2
	min, max, sum := salary[0], salary[0], 0
	for i := range salary {
		if salary[i] > max {
			max = salary[i]
		}
		if salary[i] < min {
			min = salary[i]
		}
		sum += salary[i]
	}
	return float64(sum-max-min) / float64(num)
}
