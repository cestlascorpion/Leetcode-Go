package solution

import "testing"

func TestXorOperation(t *testing.T) {
	n := 5
	start := 0
	res := XorOperation(n, start)
	exp := 8
	if res != exp {
		t.Fatalf("exp %v res %v\n", exp, res)
	}
}
