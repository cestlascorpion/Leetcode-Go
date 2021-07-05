package solution

func FindKthPositive(arr []int, k int) int {
	num := 1
	res := 1
	for idx := 0; idx < len(arr); {
		if arr[idx] != num {
			k--
			res = num
		} else {
			idx++
		}
		if k == 0 {
			break
		}
		num++
	}
	if k == 0 {
		return res
	}
	return arr[len(arr)-1] + k
}
