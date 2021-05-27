package solution

import "testing"

func TestMaximum69Number(t *testing.T) {
	num := 9669
	exp := 9969
	if exp != Maximum69Number(num) {
		t.Fail()
	}
}

func TestMaximum69Number2(t *testing.T) {
	num := 9999
	exp := 9999
	if exp != Maximum69Number(num) {
		t.Fail()
	}
}

func TestMaximum69Number3(t *testing.T) {
	num := 9996
	exp := 9999
	if exp != Maximum69Number(num) {
		t.Fail()
	}
}
