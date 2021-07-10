package solution

func MinLogOperations(logs []string) int {
	res := 0
	for i := range logs {
		switch logs[i] {
		case "../":
			res--
			if res < 0 {
				res = 0
			}
		case "./":
		default:
			res++
		}
	}
	return res
}
