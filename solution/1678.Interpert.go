package solution

func Interpret(command string) string {
	res := []byte{}
	for i := 0; i < len(command); i++ {
		if command[i] == 'G' {
			res = append(res, 'G')
		} else {
			if command[i+1] == ')' {
				res = append(res, 'o')
				i++
			} else {
				res = append(res, 'a', 'l')
				i += 3
			}
		}
	}
	return string(res)
}
