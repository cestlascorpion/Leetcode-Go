package solution

import "testing"

func TestStringMatching(t *testing.T) {
	words := []string{"mass", "as", "hero", "superhero"}
	exp := []string{"as", "hero"}
	res := StringMatching(words)
	if !IsSameStrArray(exp, res) {
		t.Fail()
	}
}
