package solution

import (
	"fmt"
	"testing"
)

func TestGuessNumber(t *testing.T) {
	exp := GetGuessNum()
	res := GuessNumber(10)
	fmt.Println(exp, res)
	if exp != res {
		t.Fail()
	}
}
