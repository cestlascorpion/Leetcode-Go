package solution

func ReplaceSpace(s string) string {
	str := make([]byte, 0)
	for i := range s {
		if s[i] == ' ' {
			str = append(str, '%', '2', '0')
		}else {
			str = append(str, s[i])
		}
	}
	return string(str)
}
