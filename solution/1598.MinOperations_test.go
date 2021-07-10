package solution

import "testing"

func TestMinLogOperations(t *testing.T) {
	Logs := []string{"d1/", "d2/", "../", "d21/", "./"}
	exp := 2
	if MinLogOperations(Logs) != exp {
		t.Fail()
	}
}

func TestMinLogOperations2(t *testing.T) {
	Logs := []string{"d1/", "d2/", "./", "d3/", "../", "d31/"}
	exp := 3
	if MinLogOperations(Logs) != exp {
		t.Fail()
	}
}

func TestMinLogOperations3(t *testing.T) {
	Logs := []string{"d1/", "../", "../", "../"}
	exp := 0
	if MinLogOperations(Logs) != exp {
		t.Fail()
	}
}

func TestMinLogOperations4(t *testing.T) {
	Logs := []string{"./", "wz4/", "../", "mj2/", "../", "../", "ik0/", "il7/"}
	exp := 2
	if MinLogOperations(Logs) != exp {
		t.Fail()
	}
}
