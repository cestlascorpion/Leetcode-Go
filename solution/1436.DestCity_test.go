package solution

import "testing"

func TestDestCity(t *testing.T) {
	paths := [][]string{
		{"B", "C"},
		{"D", "B"},
		{"C", "A"},
	}
	exp := "A"
	if DestCity(paths) != exp {
		t.Fail()
	}
}

func TestDestCity2(t *testing.T) {
	paths := [][]string{
		{"B", "C"},
		{"D", "B"},
		{"C", "A"},
	}
	exp := "A"
	if DestCity2(paths) != exp {
		t.Fail()
	}
}
