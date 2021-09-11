package solution

import (
	"strconv"
)

func GetLucky(s string, k int) int {
	var num string
	for i := range s {
		num += strconv.FormatInt(int64(s[i]-'a')+1, 10)
	}
	sum := int64(0)
	for {
		for i := range num {
			sum += int64(num[i] - '0')
		}
		num = strconv.FormatInt(sum, 10)
		k--
		if k == 0 {
			return int(sum)
		}
		sum = 0
	}
}
