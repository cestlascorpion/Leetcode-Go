package solution

func IsSumEqual(firstWord string, secondWord string, targetWord string) bool {
	i, j, k := len(firstWord), len(secondWord), len(targetWord)
	o := false
	for cur := 0; cur < k; cur++ {
		a, b := uint8(0), uint8(0)
		if cur < i {
			a = firstWord[i-cur-1] - 'a'
		}
		if cur < j {
			b = secondWord[j-cur-1] - 'a'
		}
		c := a + b
		if o {
			c++
		}
		if c > 9 {
			o = true
			c = c - 10
		} else {
			o = false
		}
		if c != targetWord[k-cur-1]-'a' {
			return false
		}
	}
	return true
}
