package solution

import "testing"

func TestMaxDepth(t *testing.T) {
	if 3 != MaxDepth("(1+(2*3)+((8)/4))+1") {
		t.Fail()
	}
	if 3 != MaxDepth("(1)+((2))+(((3)))") {
		t.Fail()
	}
	if 1 != MaxDepth("1+(2*3)/(2-1)") {
		t.Fail()
	}
}
