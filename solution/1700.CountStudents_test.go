package solution

import "testing"

func TestCountStudents(t *testing.T) {
	students := []int{1, 1, 0, 0}
	sandwiches := []int{0, 1, 0, 1}
	exp := 0
	if CountStudents(students, sandwiches) != exp {
		t.Fail()
	}
}

func TestCountStudents2(t *testing.T) {
	students := []int{1, 1, 1, 0, 0, 1}
	sandwiches := []int{1, 0, 0, 0, 1, 1}
	exp := 3
	if CountStudents(students, sandwiches) != exp {
		t.Fail()
	}
}
