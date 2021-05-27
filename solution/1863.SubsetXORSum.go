package solution

func SubsetXORSum(nums []int) int {
	res := 0 // Amazing
	for i := 0; i < (1 << len(nums)); i++ {
		xor := 0
		for j := 0; j < len(nums); j++ {
			if (i>>j)&1 != 0 {
				xor = xor ^ nums[j]
			}
		}
		res += xor
	}
	return res
}

var SubsetXORSum2Result int

func subsetHelper(val, idx int, nums []int) {
	if idx == len(nums) {
		SubsetXORSum2Result += val
		return
	}
	subsetHelper(val^nums[idx], idx+1, nums)
	subsetHelper(val, idx+1, nums)
}

func SubsetXORSum2(nums []int) int {
	SubsetXORSum2Result = 0
	subsetHelper(0, 0, nums)
	return SubsetXORSum2Result
}
