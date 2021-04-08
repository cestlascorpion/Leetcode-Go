package solution

import (
	"fmt"
	"testing"
)

func TestNumberOfSteps(t *testing.T) {
	c := [3][2]int{{14, 6}, {8, 4}, {123, 12}}
	for i := 0; i < len(c); i++ {
		res := NumberOfSteps(c[i][0])
		if c[i][1] != res {
			fmt.Println(c[i][1], res)
			t.Fail()
		}
	}
}
