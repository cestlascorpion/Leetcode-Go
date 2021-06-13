package solution

import (
	"strconv"
)

func SummaryRanges(nums []int) []string {
	res := make([]string, 0)
	if len(nums) == 0 {
		return res
	}
	for cursor := 0; cursor < len(nums); {
		begin := cursor
		step := 1
		end := begin + step
		for end != len(nums) {
			if nums[end] == nums[begin]+step {
				step++
				end++
			} else {
				break
			}
		}
		if begin == end-1 {
			res = append(res, strconv.Itoa(nums[begin]))
		} else {
			res = append(res, strconv.Itoa(nums[begin])+"->"+strconv.Itoa(nums[end-1]))
		}
		cursor = end
	}
	return res
}
