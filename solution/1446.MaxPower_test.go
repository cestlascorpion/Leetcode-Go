package solution

import "testing"

func TestMaxPower(t *testing.T) {
	s := "leetcode"
	exp := 2
	if MaxPower(s) != exp {
		t.Fail()
	}
}

func TestMaxPower2(t *testing.T) {
	s := "abbcccddddeeeeedcba"
	exp := 5
	if MaxPower(s) != exp {
		t.Fail()
	}
}

func TestMaxPower3(t *testing.T) {
	s := "triplepillooooow"
	exp := 5
	if MaxPower(s) != exp {
		t.Fail()
	}
}

func TestMaxPower4(t *testing.T) {
	s := "hooraaaaaaaaaaa"
	exp := 11
	if MaxPower(s) != exp {
		t.Fail()
	}
}
