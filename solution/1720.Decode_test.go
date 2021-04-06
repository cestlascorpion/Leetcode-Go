package solution

import "testing"

func TestDecode(t *testing.T) {
	encoded := []int{1, 2, 3}
	first := 1
	res := Decode(encoded, first)
	exp := []int{1, 0, 2, 1}
	if len(res) != len(exp) {
		t.Fail()
	}
	for i, v := range exp {
		if res[i] != v {
			t.Fail()
		}
	}
}
