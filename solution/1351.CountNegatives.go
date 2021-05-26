package solution

func CountNegatives(grid [][]int) int {
	res := 0
	row := len(grid) - 1
	col := 0
	for row >= 0 && col < len(grid[0]) {
		if grid[row][col] < 0 {
			row--
		} else {
			col++
			res += len(grid) - 1 - row
		}
	}
	if row == -1 {
		res += (len(grid) - 1 - row) * (len(grid[0]) - col)
	}
	return res
}
