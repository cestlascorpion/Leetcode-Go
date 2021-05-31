package solution

import (
	"fmt"
	"testing"
)

func TestIsPowerOfFour(t *testing.T) {
	n := [3]int{16, 5, 1}
	exp := [3]bool{true, false, true}
	for i := range n {
		if IsPowerOfFour(n[i]) != exp[i] {
			fmt.Println(n[i])
			t.Fail()
		}
	}
}
