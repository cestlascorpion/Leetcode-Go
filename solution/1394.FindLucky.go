package solution

func FindLucky(arr []int) int {
	res := -1
	cnt := make(map[int]int)
	for i := range arr {
		cnt[arr[i]] ++
	}
	for k, v := range cnt {
		if k == v {
			if k > res {
				res = k
			}
		}
	}
	return res
}
