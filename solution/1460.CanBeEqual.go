package solution

func CanBeEqual(target []int, arr []int) bool {
	if len(target) != len(arr) {
		return false
	}
	m := make(map[int]int)
	for i := 0; i < len(target); i++ {
		m[target[i]] ++
		m[arr[i]] --
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
