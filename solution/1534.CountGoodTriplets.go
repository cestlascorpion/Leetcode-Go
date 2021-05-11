package solution

func valid(x, y, length int) bool {
	if x > y {
		return x-y <= length
	}
	return y-x <= length
}

func CountGoodTriplets(arr []int, a int, b int, c int) int {
	sum := 0
	for i := 0; i < len(arr)-2; i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			if !valid(arr[i], arr[j], a) {
				continue
			}
			for k := j + 1; k < len(arr); k++ {
				if valid(arr[j], arr[k], b) && valid(arr[i], arr[k], c) {
					sum++
				}
			}
		}
	}
	return sum
}
