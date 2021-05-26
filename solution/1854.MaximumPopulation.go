package solution

func MaximumPopulation(logs [][]int) int {
	record := [100 + 1]int{}
	for i := range logs {
		record[logs[i][0]-1950]++
		record[logs[i][1]-1950]--
	}
	max := 0
	cur := 0
	idx := 0
	for i := range record {
		cur = cur + record[i]
		if cur > max {
			max = cur
			idx = i
		}
	}
	return idx + 1950
}
