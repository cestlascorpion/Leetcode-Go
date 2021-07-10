package solution

func repeat(arr []int, m, k, i int) bool {
	for mdx := 0; mdx < m; mdx++ {
		var ch int
		for kdx := 0; kdx < k; kdx++ {
			idx := i + mdx + kdx*m
			if kdx == 0 {
				ch = arr[idx]
			} else {
				if arr[idx] != ch {
					return false
				}
			}
		}

	}
	return true
}

func ContainsPattern(arr []int, m int, k int) bool {
	l := len(arr)
	for i := 0; i <= l-m*k; i++ {
		if repeat(arr, m, k, i) {
			return true
		}
	}
	return false
}
