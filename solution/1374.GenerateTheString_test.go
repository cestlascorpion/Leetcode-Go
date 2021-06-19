package solution

import "testing"

func TestGenerateTheString(t *testing.T) {
	m := make(map[uint8]int)
	s := GenerateTheString(10)
	for i := range s {
		m[s[i]]++
	}
	for _, v := range m {
		if v%2 == 0 {
			t.Fail()
		}
	}
}
