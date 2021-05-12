package solution

import "testing"

func TestKthToLast(t *testing.T) {
	list := constructList()
	exp := 4
	if KthToLast(list, 2) != exp {
		t.Fail()
	}
}
