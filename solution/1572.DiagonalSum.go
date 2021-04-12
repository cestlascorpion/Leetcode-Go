package solution

func DiagonalSum(mat [][]int) int {
	res := 0
	l := len(mat)
	for i, j := 0, 0; i < l/2; i++ {
		res += mat[i][j] + mat[l-i-1][j] + mat[i][l-j-1] + mat[l-i-1][l-j-1]
		j++
	}
	if (l & 1) == 1 {
		res += mat[l/2][l/2]
	}
	return res
}
