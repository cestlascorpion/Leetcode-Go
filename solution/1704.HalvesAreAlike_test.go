package solution

import "testing"

func TestHalvesAreAlike(t *testing.T) {
	s := "book"
	exp := true
	if HalvesAreAlike(s) != exp {
		t.Fail()
	}
}

func TestHalvesAreAlike2(t *testing.T) {
	s := "textbook"
	exp := false
	if HalvesAreAlike(s) != exp {
		t.Fail()
	}
}

func TestHalvesAreAlike3(t *testing.T) {
	s := "MerryChristmas"
	exp := false
	if HalvesAreAlike(s) != exp {
		t.Fail()
	}
}

func TestHalvesAreAlike4(t *testing.T) {
	s := "AbCdEfGh"
	exp := true
	if HalvesAreAlike(s) != exp {
		t.Fail()
	}
}
