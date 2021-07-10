package solution

import (
	"fmt"
	"testing"
)

func TestModifyString(t *testing.T) {
	s := "ubv?w"
	res := ModifyString(s)
	fmt.Println(res)
}

func TestModifyString2(t *testing.T) {
	s := "j?qg??b"
	res := ModifyString(s)
	fmt.Println(res)
}

func TestModifyString3(t *testing.T) {
	s := "??yw?ipkj?"
	res := ModifyString(s)
	fmt.Println(res)
}
