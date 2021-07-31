package solution

import "testing"

func TestMakeEqual(t *testing.T) {
	words:=[]string{"abc","aabc","bc"}
	if !MakeEqual(words) {
		t.Fail()
	}
}