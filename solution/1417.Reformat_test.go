package solution

import (
	"fmt"
	"testing"
)

func TestReformat(t *testing.T) {
	s := "a0b1c2"
	fmt.Println(Reformat(s))
}

func TestReformat2(t *testing.T) {
	s := "covid2019"
	fmt.Println(Reformat(s))
}
