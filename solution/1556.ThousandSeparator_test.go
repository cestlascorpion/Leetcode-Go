package solution

import (
	"fmt"
	"testing"
)

func TestThousandSeparator(t *testing.T) {
	n := 123456789
	exp := "123.456.789"
	res := ThousandSeparator(n)
	if res != exp {
		fmt.Println(res)
		t.Fail()
	}
}
