package solution

func maximumWealth(accounts [][]int) int {
	ret := 0
	for i := range accounts {
		temp := 0
		for j := range accounts[i] {
			temp += accounts[i][j]
		}
		if temp > ret {
			ret = temp
		}
	}
	return ret
}
