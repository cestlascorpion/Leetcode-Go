package solution

func PrefixesDivBy5(nums []int) []bool {
	res := make([]bool, len(nums))
	x := 0
	for i, v := range nums {
		x = (x<<1 | v) % 5
		res[i] = x == 0
	}
	return res
}
