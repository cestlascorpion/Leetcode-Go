package solution

import "testing"

func TestMaxAscendingSum(t *testing.T) {
	nums := []int{10, 20, 30, 5, 10, 50}
	exp := 65
	if MaxAscendingSum(nums) != exp {
		t.Fail()
	}
}

func TestMaxAscendingSum2(t *testing.T) {
	nums := []int{10,20,30,40,50}
	exp := 150
	if MaxAscendingSum(nums) != exp {
		t.Fail()
	}
}


func TestMaxAscendingSum3(t *testing.T) {
	nums := []int{12,17,15,13,10,11,12}
	exp := 33
	if MaxAscendingSum(nums) != exp {
		t.Fail()
	}
}

func TestMaxAscendingSum4(t *testing.T) {
	nums := []int{100,10,1}
	exp := 100
	if MaxAscendingSum(nums) != exp {
		t.Fail()
	}
}
