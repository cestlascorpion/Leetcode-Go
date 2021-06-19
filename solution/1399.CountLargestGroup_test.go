package solution

import "testing"

func TestCountGroup(t *testing.T) {
	num := 10
	exp := 1
	if CountGroup(num) != exp {
		t.Fail()
	}
}

func TestCountLargestGroup(t *testing.T) {
	n := 13
	exp := 4
	if CountLargestGroup(n) != exp {
		t.Fail()
	}
}

func TestCountLargestGroup2(t *testing.T) {
	n := 15
	exp := 6
	if CountLargestGroup(n) != exp {
		t.Fail()
	}
}

func TestCountLargestGroup3(t *testing.T) {
	n := 24
	exp := 5
	if CountLargestGroup(n) != exp {
		t.Fail()
	}
}
