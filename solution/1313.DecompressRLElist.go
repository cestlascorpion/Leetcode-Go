package solution

func DecompressRLElist(nums []int) []int {
	res := make([]int, 0)
	for i := 1; i < len(nums); i += 2 {
		for j := 0; j < nums[i-1]; j++ {
			res = append(res, nums[i])
		}
	}
	return res
}
