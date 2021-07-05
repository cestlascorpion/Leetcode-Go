package solution

func ThousandSeparator(n int) string {
	if n == 0 {
		return "0"
	}
	res := make([]byte, 0)
	cnt := 0
	for n > 0 {
		mod := n % 10
		n = n / 10
		cnt++
		res = append(res, byte(mod+'0'))
		if cnt == 3 && n != 0 {
			res = append(res, '.')
			cnt = 0
		}
	}
	for i, j := 0, len(res)-1; i < j; {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return string(res)
}
