package solution

import (
	"math"
)

func Maximum69Number(num int) int {
	last := -1
	for i, n := 0, num; n > 0; i++ {
		if (n % 10) == 6 {
			last = i
		}
		n = n / 10
	}
	if last != -1 {
		num += int(math.Pow(10, float64(last))) * 3
	}
	return num
}
