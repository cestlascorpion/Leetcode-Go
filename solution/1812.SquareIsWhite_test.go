package solution

import "testing"

func TestSquareIsWhite(t *testing.T) {
	coordinates := []string{"a1", "h3", "c7"}
	exp := []bool{false, true, false}
	for i := 0; i < len(coordinates); i++ {
		if exp[i] != SquareIsWhite(coordinates[i]) {
			t.Fail()
		}
	}
}
