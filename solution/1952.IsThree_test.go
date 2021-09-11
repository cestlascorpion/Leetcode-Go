package solution

import "testing"

func TestIsThree(t *testing.T) {
	n := 2
	exp := false
	if IsThree(n) != exp {
		t.Failed()
	}
}
