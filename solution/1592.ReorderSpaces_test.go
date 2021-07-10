package solution

import "testing"

func TestReorderSpaces(t *testing.T) {
	text := "  this   is  a sentence "
	exp := "this   is   a   sentence"
	res := ReorderSpaces(text)
	if exp != res {
		t.Fail()
	}
}

func TestReorderSpaces2(t *testing.T) {
	text := " practice   makes   perfect"
	exp := "practice   makes   perfect "
	res := ReorderSpaces(text)
	if exp != res {
		t.Fail()
	}
}

func TestReorderSpaces3(t *testing.T) {
	text := "hello   world"
	exp := "hello   world"
	res := ReorderSpaces(text)
	if exp != res {
		t.Fail()
	}
}

func TestReorderSpaces4(t *testing.T) {
	text := "  walks  udp package   into  bar a"
	exp := "walks  udp  package  into  bar  a "
	res := ReorderSpaces(text)
	if exp != res {
		t.Fail()
	}
}

func TestReorderSpaces5(t *testing.T) {
	text := " walks"
	exp := "walks "
	res := ReorderSpaces(text)
	if exp != res {
		t.Fail()
	}
}
