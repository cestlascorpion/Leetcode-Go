package solution

import (
	"testing"
)

func Tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	a, b, c := 0, 1, 1
	for i := 3; i <= n; i++ {
		a, b, c = b, c, a+b+c
	}
	return c
}

func TestTribonacci(t *testing.T) {
	n := 4
	exp := 4
	if Tribonacci(n) != exp {
		t.Fail()
	}
}
