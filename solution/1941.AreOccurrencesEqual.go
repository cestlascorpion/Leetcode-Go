package solution

func AreOccurrencesEqual(s string) bool {
	letter := [26]int{0}
	for i := range s {
		letter[s[i]-'a']++
	}
	cnt := 0
	for i := range letter {
		if letter[i] > 0 {
			if cnt == 0 {
				cnt = letter[i]
			} else {
				if cnt != letter[i] {
					return false
				}
			}
		}
	}
	return true
}
