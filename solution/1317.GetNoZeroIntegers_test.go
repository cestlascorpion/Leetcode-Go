package solution

import (
	"fmt"
	"testing"
)

func TestGetNoZeroIntegers(t *testing.T) {
	n := 1000
	res := GetNoZeroIntegers(n)
	fmt.Println(res)
}
