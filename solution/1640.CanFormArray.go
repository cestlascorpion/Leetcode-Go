package solution

func CanFormArray(arr []int, pieces [][]int) bool {
	index := make(map[int]int)
	for i := range pieces {
		index[pieces[i][0]] = i
	}
	for i := 0; i < len(arr); {
		if v, ok := index[arr[i]]; !ok {
			return false
		} else {
			for j := range pieces[v] {
				if arr[i+j] != pieces[v][j] {
					return false
				}
			}
			i += len(pieces[v])
		}
	}
	return true
}
