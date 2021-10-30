package solution

func ReversePrefix(word string, ch byte) string {
	i, j := 0, -1
	for i := range word {
		if word[i] == ch {
			j = i
			break
		}
	}
	if j == -1 || i == j {
		return word
	}
	ans := []byte(word)
	for i < j {
		ans[i], ans[j] = ans[j], ans[i]
		i++
		j--
	}
	return string(ans)
}
