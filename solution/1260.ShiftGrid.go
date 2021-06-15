package solution

func ShiftGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	k = m*n - k%(m*n)

	data := make([]int, 0, m*n)
	for i := range grid {
		data = append(data, grid[i]...)
	}
	data = append(data[k:], data[:k]...)

	res := make([][]int, 0)
	for i, skip := 0, 0; i < m; i++ {
		res = append(res, data[skip:skip+n])
		skip += n
	}
	return res
}
