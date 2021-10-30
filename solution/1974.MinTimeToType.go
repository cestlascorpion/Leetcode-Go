package solution

func MinTimeToType(word string) int {
	cost := 0
	ch := uint8('a')
	for i := range word {
		c1, c2 := (word[i]+26-ch)%26, (ch+26-word[i])%26
		if c1 < c2 {
			cost += int(c1)
		} else {
			cost += int(c2)
		}
		ch = word[i]
	}
	return cost + len(word)
}
