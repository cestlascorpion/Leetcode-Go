package solution

import (
	"fmt"
	"testing"
)

func TestFindNumbers(t *testing.T) {
	nums := []int{12, 345, 2, 6, 7896}
	res := FindNumbers(nums)
	exp := 2
	if res != exp {
		fmt.Println(res, exp)
		t.Fail()
	}
}

func TestFindNumbers2(t *testing.T) {
	nums := []int{100000}
	res := FindNumbers(nums)
	exp := 1
	if res != exp {
		fmt.Println(res, exp)
		t.Fail()
	}
}
