package solution

import "testing"

func TestCountTriples(t *testing.T) {
	n := 5
	exp := 2
	if CountTriples(n) != exp {
		t.Failed()
	}
}

func TestCountTriples2(t *testing.T) {
	n := 10
	exp := 4
	if CountTriples(n) != exp {
		t.Failed()
	}
}

func TestCountTriples3(t *testing.T) {
	n := 5
	exp := 2
	if CountTriples2(n) != exp {
		t.Failed()
	}
}

func TestCountTriples4(t *testing.T) {
	n := 10
	exp := 4
	if CountTriples2(n) != exp {
		t.Failed()
	}
}