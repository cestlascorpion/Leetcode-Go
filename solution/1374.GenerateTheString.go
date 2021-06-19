package solution

func GenerateTheString(n int) string {
	res := make([]uint8, n)
	for i := range res {
		res[i] = 'a'
	}
	if n%2 == 0 {
		res[n-1] = 'b'
	}
	return string(res)
}
