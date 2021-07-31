package solution

func IsCovered(ranges [][]int, left int, right int) bool {
	n := right - left + 1
	arr := make([]int, n)
	for i := range ranges {
		for j := ranges[i][0]; j <= ranges[i][1]; j++ {
			idx := j - left
			if idx >= 0 && idx < n {
				arr[idx]++
			}
		}
	}
	for i := range arr {
		if arr[i] == 0 {
			return false
		}
	}
	return true
}
