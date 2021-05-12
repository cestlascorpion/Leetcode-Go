package solution

import (
	"fmt"
	"testing"
)

func TestConstructorOrderedStream(t *testing.T) {
	o:= ConstructorOrderedStream(5)
	fmt.Println(o.Insert(3,"c"))
	fmt.Println(o.Insert(1,"a"))
	fmt.Println(o.Insert(2,"b"))
	fmt.Println(o.Insert(5,"e"))
	fmt.Println(o.Insert(4,"d"))
}