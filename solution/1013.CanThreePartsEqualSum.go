package solution

func CanThreePartsEqualSum(arr []int) bool {
	sum := 0
	for idx := range arr {
		sum += arr[idx]
	}
	if sum%3 != 0 {
		return false
	}
	sum /= 3
	tmp := 0
	cnt := 0
	for idx := range arr {
		tmp += arr[idx]
		if tmp == sum {
			cnt++
			if cnt == 2 && idx != len(arr)-1 {
				return true
			}
			tmp = 0
		}
	}
	return false
}
