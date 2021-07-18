package solution

func ReformatNumber(number string) string {
	nums := make([]byte, 0)
	for i := range number {
		switch number[i] {
		case ' ':
		case '-':
		default:
			nums = append(nums, number[i])
		}
	}
	res := make([]byte, 0)
	for i := 0; i < len(nums); i += 3 {
		if i+4 < len(nums) {
			res = append(res, nums[i], nums[i+1], nums[i+2], '-')
		} else {
			if i+4 == len(nums) {
				res = append(res, nums[i], nums[i+1], '-', nums[i+2], nums[i+3])
			} else if i+3 == len(nums) {
				res = append(res, nums[i], nums[i+1], nums[i+2])
			} else {
				res = append(res, nums[i:]...)
			}
			break
		}
	}
	return string(res)
}
