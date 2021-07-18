package solution

func Decrypt(code []int, k int) []int {
	res := make([]int, len(code))
	if k == 0 {
		return res
	}

	if k > 0 {
		for i := 1; i <= k; i++ {
			res[0] += code[i]
		}
	} else {
		for i := 1; i <= -k; i++ {
			res[0] += code[len(code)-i]
		}
	}
	for i := 1; i < len(res); i++ {
		if k > 0 {
			res[i] = res[i-1] - code[i] + code[(i+k)%len(code)]
		} else {
			res[i] = res[i-1] - code[(i+k-1+len(code))%len(code)] + code[i-1]
		}
	}
	return res
}
