package solution

func Reformat(s string) string {
	num, ch := make([]byte, 0), make([]byte, 0)
	for i := range s {
		if s[i] >= '0' && s[i] <= '9' {
			num = append(num, s[i])
		} else {
			ch = append(ch, s[i])
		}
	}
	if len(num) > len(ch)+1 || len(ch) > len(num)+1 {
		return ""
	}
	res := make([]byte, 0, len(s))
	i, j := 0, 0
	for i != len(num) && j != len(ch) {
		res = append(res, num[i])
		res = append(res, ch[j])
		i++
		j++
	}
	if i != len(num) {
		res = append(res, num[i])

	}
	if j != len(ch) {
		res = append([]byte{ch[j]}, res...)
	}
	return string(res)
}
