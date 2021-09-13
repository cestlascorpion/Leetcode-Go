package solution

import "fmt"

func MakeFancyString(s string) string {
	str := make([]byte, 0)
	str = append(str, s[0])
	ch := s[0]
	cnt := 1
	for i := 1; i < len(s); i++ {
		if ch == s[i] {
			if cnt < 2 {
				str = append(str, s[i])
				cnt++
			}
		} else {
			str = append(str, s[i])
			ch = s[i]
			cnt = 1
		}
	}
	fmt.Println(string(str))
	return string(str)
}
