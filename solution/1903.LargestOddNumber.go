package solution

func LargestOddNumber(num string) string {
	last := -1
	for i := len(num) - 1; i >= 0; i-- {
		if (num[i]-'0')%2 != 0 {
			last = i
			break
		}
	}
	if last == -1 {
		return ""
	}
	return num[0 : last+1]
}
