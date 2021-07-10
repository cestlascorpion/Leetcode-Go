package solution

func MostVisited(n int, rounds []int) []int {
	s := rounds[0]
	e := rounds[len(rounds)-1]
	res := make([]int, 0)
	if e >= s {
		for i := s; i <= e; i++ {
			res = append(res, i)
		}
	} else {
		for i := 1; i <= e; i++ {
			res = append(res, i)
		}
		for i := s; i <= n; i++ {
			res = append(res, i)
		}
	}
	return res
}
