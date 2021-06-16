package solution

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func AllCellsDistOrder(rows int, cols int, rCenter int, cCenter int) [][]int {
	maxDist := max(rCenter, rows-1-rCenter) + max(cCenter, cols-1-cCenter)
	buckets := make([][][]int, maxDist+1)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			dist := abs(i-rCenter) + abs(j-cCenter)
			buckets[dist] = append(buckets[dist], []int{i, j})
		}
	}

	ans := make([][]int, 0, rows*cols)
	for _, bucket := range buckets {
		ans = append(ans, bucket...)
	}
	return ans
}
