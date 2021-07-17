package solution

import "testing"

func TestSlowestKey(t *testing.T) {
	releaseTimes := []int{12, 23, 36, 46, 62}
	keysPressed := "spuda"
	exp := byte('a')
	if SlowestKey(releaseTimes, keysPressed) != exp {
		t.Fail()
	}
}
