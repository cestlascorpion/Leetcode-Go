package solution

import "testing"

func TestCountMatches(t *testing.T) {
	items := [][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "lenovo"}, {"phone", "gold", "iphone"}}
	ruleKey := "color"
	ruleValue := "silver"
	res := CountMatches(items, ruleKey, ruleValue)
	exp := 1
	if res != exp {
		t.Fail()
	}
}
