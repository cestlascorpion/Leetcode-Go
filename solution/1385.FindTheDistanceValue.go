package solution

func FindTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	res := 0
	for i := range arr1 {
		for j := range arr2 {
			var delta int
			if arr1[i] > arr2[j] {
				delta = arr1[i] - arr2[j]
			} else {
				delta = arr2[j] - arr1[i]
			}
			if delta <= d {
				break
			}
			if j == len(arr2)-1 {
				res++
			}
		}
	}
	return res
}
