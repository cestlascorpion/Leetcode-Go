package solution

func CheckIfPangram(sentence string) bool {
	dict := [26]int{}
	for i := range sentence {
		index := sentence[i] - 'a'
		dict[index]++
	}
	for i := range dict {
		if dict[i] == 0 {
			return false
		}
	}
	return true
}
