package solution

import "testing"

func TestNumWaterBottles(t *testing.T) {
	numBottles := 9
	numExchange := 3
	exp := 13
	if NumWaterBottles(numBottles, numExchange) != exp {
		t.Fail()
	}
}

func TestNumWaterBottles2(t *testing.T) {
	numBottles := 15
	numExchange := 4
	exp := 19
	if NumWaterBottles(numBottles, numExchange) != exp {
		t.Fail()
	}
}

func TestNumWaterBottles3(t *testing.T) {
	numBottles := 5
	numExchange := 5
	exp := 6
	if NumWaterBottles(numBottles, numExchange) != exp {
		t.Fail()
	}
}

func TestNumWaterBottles4(t *testing.T) {
	numBottles := 2
	numExchange := 3
	exp := 2
	if NumWaterBottles(numBottles, numExchange) != exp {
		t.Fail()
	}
}
