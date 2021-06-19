package solution

func CheckIfExist(arr []int) bool {
	m := make(map[int]int)
	for i := range arr {
		if _, two := m[arr[i]*2]; two {
			return true
		}
		if _, half := m[arr[i]/2]; arr[i]%2 == 0 && half {
			return true
		}
		m[arr[i]]++
	}
	return false
}
