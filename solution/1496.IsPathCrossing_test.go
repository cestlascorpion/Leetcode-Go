package solution

import "testing"

func TestIsPathCrossing(t *testing.T) {
	path:="NESWW"
	if !IsPathCrossing(path) {
		t.Fail()
	}
}