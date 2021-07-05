package solution

const Delta = 'a' - 'A'

func MakeGood(s string) string {
	str := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if len(str) == 0 {
			str = append(str, s[i])
		} else {
			if str[len(str)-1] == s[i]+Delta || str[len(str)-1] == s[i]-Delta {
				str = str[0 : len(str)-1]
			} else {
				str = append(str, s[i])
			}
		}
	}
	return string(str)
}
