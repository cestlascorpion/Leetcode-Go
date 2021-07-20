package solution

import "testing"

func TestCountBalls(t *testing.T) {
	low:=1
	high:=10
	exp:=2
	if CountBalls(low,high)!=exp {
		t.Fail()
	}
}