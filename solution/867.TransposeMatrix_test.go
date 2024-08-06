package solution

import "testing"

func Transpose(matrix [][]int) [][]int {
	row, col := len(matrix), len(matrix[0])
	transposed := make([][]int, col)
	for i := 0; i < col; i++ {
		transposed[i] = make([]int, row)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			transposed[j][i] = matrix[i][j]
		}
	}
	return transposed
}

func TestTranspose(t *testing.T) {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	target := [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}
	result := Transpose(matrix)
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[i]); j++ {
			if result[i][j] != target[i][j] {
				t.Fatalf("867. Transpose Matrix: case 1 failed.")
			}
		}
	}
	t.Log("867. Transpose Matrix: case 1 succeeded.")
}
