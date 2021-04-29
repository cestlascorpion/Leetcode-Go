package solution

import (
	"fmt"
	"testing"
)

func TestReverseLeftWords(t *testing.T) {
	s := "abcdefg"
	k := 2
	exp := "cdefgab"
	res := ReverseLeftWords(s, k)
	if exp != res {
		fmt.Println(res, exp)
		t.Fail()
	}
}
