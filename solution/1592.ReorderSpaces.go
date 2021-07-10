package solution

func isBegin(text *string, i int) bool {
	if i == 0 {
		return true
	}
	if (*text)[i-1] == ' ' {
		return true
	}
	return false
}

func isEnd(text *string, i int) bool {
	if i == len(*text)-1 {
		return true
	}
	if (*text)[i+1] == ' ' {
		return true
	}
	return false
}

func ReorderSpaces(text string) string {
	words := make([][2]int, 0)
	space := 0
	for i := 0; i < len(text); i++ {
		if text[i] == ' ' {
			space++
			continue
		}
		if isBegin(&text, i) {
			words = append(words, [2]int{i, i})
		}
		if isEnd(&text, i) {
			words[len(words)-1][1] = i
		}
	}
	var (
		div int
		rst int
	)
	if len(words) == 1 {
		div = 0
		rst = space
	} else {
		div = space / (len(words) - 1)
		rst = space % (len(words) - 1)
	}
	res := make([]byte, 0, len(text))
	for w := range words {
		for i := words[w][0]; i <= words[w][1]; i++ {
			res = append(res, text[i])
		}
		if w == len(words)-1 {
			break
		}
		for i := 0; i < div; i++ {
			res = append(res, ' ')
		}
	}
	for i := 0; i < rst; i++ {
		res = append(res, ' ')
	}

	return string(res)
}
