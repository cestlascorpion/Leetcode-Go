package solution

import (
	"testing"
)

func BalancedStringSplit(s string) int {
	n := 0
	x := 0
	for i := range s {
		if s[i] == 'L' {
			x++
		} else {
			x--
		}
		if x == 0 {
			n++
		}
	}
	return n
}

func TestBalancedStringSplit(t *testing.T) {
	s := "RLRRLLRLRL"
	exp := 4
	if BalancedStringSplit(s) != exp {
		t.Fail()
	}

	s = "RLLLLRRRLR"
	exp = 3
	if BalancedStringSplit(s) != exp {
		t.Fail()
	}

	s = "LLLLRRRR"
	exp = 1
	if BalancedStringSplit(s) != exp {
		t.Fail()
	}
}
