package solution

import (
	"fmt"
	"testing"
)

func cnIsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	z := x
	y := 0
	for x > 0 {
		y = y*10 + x%10
		x = x / 10
	}
	return z == y
}

func TestIsPalindrome(t *testing.T) {
	fmt.Println(cnIsPalindrome(0))
	fmt.Println(cnIsPalindrome(121))
	fmt.Println(cnIsPalindrome(-121))
	fmt.Println(cnIsPalindrome(123))
}
