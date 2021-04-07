package solution

import (
	"fmt"
	"testing"
)

func TestInterpret(t *testing.T) {
	command := "(al)G(al)()()G"
	res := Interpret(command)
	exp := "alGalooG"
	if exp != res {
		fmt.Println(res)
		t.Fail()
	}
}
