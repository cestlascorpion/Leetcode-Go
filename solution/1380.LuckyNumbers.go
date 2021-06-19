package solution

func LuckyNumbers(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	num := make([]int, m+n)
	for i := range matrix {
		for j := range matrix[i] {
			if num[i] == 0 || num[i] > matrix[i][j] {
				num[i] = matrix[i][j]
			}
			if num[m+j] == 0 || num[m+j] < matrix[i][j] {
				num[m+j] = matrix[i][j]
			}
		}
	}
	res := make([]int, 0)
	set := make(map[int]struct{})
	for i := range num {
		if _, ok := set[num[i]]; ok {
			res = append(res, num[i])
		} else {
			set[num[i]] = struct{}{}
		}
	}
	return res
}
