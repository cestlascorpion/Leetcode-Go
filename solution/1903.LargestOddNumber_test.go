package solution

import "testing"

func TestLargestOddNumber(t *testing.T) {
	num := "52"
	exp := "5"
	if LargestOddNumber(num) != exp {
		t.Fail()
	}
}

func TestLargestOddNumber2(t *testing.T) {
	num := "2"
	exp := ""
	if LargestOddNumber(num) != exp {
		t.Fail()
	}
}
