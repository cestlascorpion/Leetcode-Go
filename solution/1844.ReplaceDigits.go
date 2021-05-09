package solution

func shift(c, x uint8) uint8 {
	return c + (x - '0')
}

func ReplaceDigits(s string) string {
	arr := []uint8(s)
	for i := 1; i < len(arr); i += 2 {
		arr[i] = shift(arr[i], arr[i-1])
	}
	return string(arr)
}
