package solution

import (
	"fmt"
	"testing"
)

func TestReplaceDigits(t *testing.T) {
	s := "a1c1e1"
	exp := "abcdef"
	res := ReplaceDigits(s)
	if len(res) != len(exp) {
		t.Fail()
	}
	fmt.Println(res, exp)
	for i, v := range exp {
		if res[i] != uint8(v) {
			t.Fail()
		}
	}
}
