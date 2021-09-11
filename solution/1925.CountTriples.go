package solution

import "math"

func CountTriples(n int) int {
	cnt := 0
	max := n * n
	for a := 1; a <= n; a++ {
		for b := a; b <= n; b++ {
			s := a*a + b*b
			if s > max {
				break
			}
			r := math.Sqrt(float64(s))
			if math.Trunc(r) == r {
				cnt++
				if a != b {
					cnt++
				}
			}
		}
	}
	return cnt
}

func CountTriples2(n int) int {
	all := make(map[int]struct{}, 0)
	for i := 1; i <= n; i++ {
		all[i*i] = struct{}{}
	}
	cnt := 0

	for a := 1; a <= n; a++ {
		for b := a; b <= n; b++ {
			s := a*a + b*b
			if _, ok := all[s]; ok {
				cnt++
				if a != b {
					cnt++
				}
			}
		}
	}
	return cnt
}
