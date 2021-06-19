package solution

func CountGroup(num int) int {
	res := 0
	for num > 0 {
		res += num % 10
		num /= 10
	}
	return res
}

func CountLargestGroup(n int) int {
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		m[CountGroup(i+1)]++
	}
	cnt := 0
	max := -1
	for _, v := range m {
		if v > max {
			max = v
			cnt = 1
		} else if v == max {
			cnt++
		} else {
			continue
		}
	}
	return cnt
}
