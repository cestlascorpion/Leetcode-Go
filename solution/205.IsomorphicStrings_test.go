package solution

import "testing"

func IsomorphicStrings(s1 string, s2 string) bool {
	d1 := make(map[byte]byte)
	d2 := make(map[byte]byte)
	for i := 0; i < len(s1); i++ {
		if v, ok := d1[s1[i]]; ok {
			if v != s2[i] {
				return false
			}
		} else {
			d1[s1[i]] = s2[i]
		}
		if v, ok := d2[s2[i]]; ok {
			if v != s1[i] {
				return false
			}
		} else {
			d2[s2[i]] = s1[i]
		}
	}
	return true
}

func TestIsomorphicStrings(t *testing.T) {
	s1, s2 := "egg", "add"
	if IsomorphicStrings(s1, s2) {
		t.Log("205. Isomorphic Strings: case 1 succeeded.")
	} else {
		t.Error("205. Isomorphic Strings: case 1 failed.")
	}
	s1, s2 = "foo", "bar"
	if !IsomorphicStrings(s1, s2) {
		t.Log("205. Isomorphic Strings: case 2 succeeded.")
	} else {
		t.Error("205. Isomorphic Strings: case 2 failed.")
	}
	s1, s2 = "paper", "title"
	if IsomorphicStrings(s1, s2) {
		t.Log("205. Isomorphic Strings: case 3 succeeded.")
	} else {
		t.Error("205. Isomorphic Strings: case 3 failed.")
	}
	s1, s2 = "badc", "baba"
	if !IsomorphicStrings(s1, s2) {
		t.Log("205. Isomorphic Strings: case 4 succeeded.")
	} else {
		t.Error("205. Isomorphic Strings: case 4 failed.")
	}
}
