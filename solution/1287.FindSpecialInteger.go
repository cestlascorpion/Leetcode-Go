package solution

func FindSpecialInteger(arr []int) int {
	for i, l := 0, len(arr)/4; i < len(arr)-l; i++ {
		if arr[i] == arr[i+l] {
			return arr[i]
		}
	}
	return arr[0]
}
