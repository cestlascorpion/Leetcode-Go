package solution

import (
	"strconv"
)

func str2Integer(s string) uint8 {
	idx, err := strconv.Atoi(s)
	if err != nil {
		panic("fuck")
	}
	return uint8(idx)
}

func FreqAlphabets(s string) string {
	res := make([]uint8, 0, len(s))
	l := len(s)
	for i := 0; i < l; {
		if i+2 < l && s[i+2] == '#' {
			res = append(res, str2Integer(s[i:i+2])-1+'a')
			i += 3
		} else {
			res = append(res, s[i]-'1'+'a')
			i++
		}

	}
	return string(res)
}
