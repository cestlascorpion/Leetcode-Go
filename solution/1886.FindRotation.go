package solution

func FindRotation(mat [][]int, target [][]int) bool {
	n := len(mat)
	v := [4]bool{true, true, true, true}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if v[0] && mat[i][j] != target[i][j] {
				v[0] = false
			}
			if v[1] && mat[i][j] != target[n-i-1][n-j-1] {
				v[1] = false
			}
			if v[2] && mat[i][j] != target[j][n-i-1] {
				v[2] = false
			}
			if v[3] && mat[i][j] != target[n-1-j][i] {
				v[3] = false
			}
		}
	}
	for i := range v {
		if v[i] {
			return true
		}
	}
	return false
}
