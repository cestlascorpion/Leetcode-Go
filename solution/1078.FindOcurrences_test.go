package solution

import "testing"

func TestFindOcurrences(t *testing.T) {
	text := "alice is a good girl she is a good student"
	exp := []string{"girl", "student"}
	res := FindOcurrences(text, "a", "good")
	if !IsSameStrArray(exp, res) {
		t.Fail()
	}
}
