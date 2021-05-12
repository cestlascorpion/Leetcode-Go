package solution

import "testing"

func TestTruncateSentence(t *testing.T) {
	s := "Hello how are you Contestant"
	k := 4
	exp := "Hello how are you"
	res := TruncateSentence(s, k)
	if len(exp) != len(res) {
		t.Fail()
	}
	for i := range exp {
		if exp[i] != res[i] {
			t.Fail()
		}
	}
}

func TestTruncateSentence2(t *testing.T) {
	s := "Hello how are you"
	k := 4
	exp := "Hello how are you"
	res := TruncateSentence(s, k)
	if len(exp) != len(res) {
		t.Fail()
	}
	for i := range exp {
		if exp[i] != res[i] {
			t.Fail()
		}
	}
}

func TestTruncateSentence3(t *testing.T) {
	s := "Hello"
	k := 1
	exp := "Hello"
	res := TruncateSentence(s, k)
	if len(exp) != len(res) {
		t.Fail()
	}
	for i := range exp {
		if exp[i] != res[i] {
			t.Fail()
		}
	}
}
