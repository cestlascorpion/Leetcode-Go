package solution

func FindShortestSubArray(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	max := 0
	for _, v := range m {
		if v > max {
			max = v
		}
	}
	num := make([]int, 0)
	for k, v := range m {
		if v == max {
			num = append(num, k)
		}
	}
	min := len(nums)
	for _, n := range num {
		left, right := 0, len(nums)-1
		for nums[left] != n {
			left++
		}
		for nums[right] != n {
			right--
		}
		if right-left+1 < min {
			min = right - left + 1
		}
	}
	return min
}

func FindShortestSubArray2(nums []int) int {
	m := make(map[int][]int)
	for i, num := range nums {
		if _, ok := m[num]; !ok {
			m[num] = []int{i, i, 1}
		} else {
			m[num][1] = i
			m[num][2]++
		}
	}
	max := 0
	for _, v := range m {
		if v[2] > max {
			max = v[2]
		}
	}
	min := len(nums)
	for _, v := range m {
		if v[2] == max {
			if v[1]-v[0]+1 < min {
				min = v[1] - v[0] + 1
			}
		}
	}
	return min
}
