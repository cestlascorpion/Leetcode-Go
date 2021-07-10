package solution

func isSpecial(mat [][]int, r, c, i, j int) bool {
	for row := 0; row < r; row++ {
		if row != i && mat[row][j] != 0 {
			return false
		}
	}
	for col := 0; col < c; col++ {
		if col != j && mat[i][col] != 0 {
			return false
		}
	}
	return true
}

func NumSpecial(mat [][]int) int {
	res := 0
	row := len(mat)
	col := len(mat[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if mat[i][j] == 1 && isSpecial(mat, row, col, i, j) {
				res++
			}
		}
	}
	return res
}
