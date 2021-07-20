package solution

func SumOfUnique(nums []int) int {
	dict := [100]int{}
	for i := range nums {
		dict[nums[i]-1]++
	}
	sum := 0
	for i := range dict {
		if dict[i] == 1 {
			sum += i+1
		}
	}
	return sum
}
