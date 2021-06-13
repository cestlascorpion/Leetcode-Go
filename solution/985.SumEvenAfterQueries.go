package solution

func SumEvenAfterQueries(nums []int, queries [][]int) []int {
	res := make([]int, len(queries), len(queries))
	sum := 0
	for i := range nums {
		if (nums[i] % 2) == 0 {
			sum += nums[i]
		}
	}
	for i := range queries {
		val := queries[i][0]
		idx := queries[i][1]
		temp := nums[idx] + val
		if nums[idx]%2 == 0 {
			if temp%2 == 0 {
				sum += val
			} else {
				sum -= nums[idx]
			}
		} else {
			if temp%2 == 0 {
				sum += temp
			}
		}
		res[i] = sum
		nums[idx] = temp
	}
	return res
}
