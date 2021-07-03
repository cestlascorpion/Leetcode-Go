package solution

const (
	Push = "Push"
	Pop  = "Pop"
)

func BuildArray(target []int, n int) []string {
	res := make([]string, 0)
	num := 1
	for idx := range target {
		for num != target[idx] {
			res = append(res, Push, Pop)
			num++
		}
		res = append(res, Push)
		num++
	}
	return res
}
