package solution

func FindJudge(n int, trust [][]int) int {
	analyze := make([][2]int, n)
	for i := range trust {
		analyze[trust[i][0]-1][0]++
		analyze[trust[i][1]-1][1]++
	}
	for i := range analyze {
		if analyze[i][0] == 0 && analyze[i][1] == n-1 {
			return i + 1
		}
	}
	return -1
}
