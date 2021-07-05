package solution

func ThreeConsecutiveOdds(arr []int) bool {
	cnt := 0
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 != 0 {
			cnt++
		} else {
			cnt = 0
		}
		if cnt >= 3 {
			return true
		}
	}
	return false
}
