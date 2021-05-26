package solution

import "testing"

func TestMaximumPopulation(t *testing.T) {
	logs := [][]int{{1993, 1999}, {2000, 2010}}
	exp := 1993
	res := MaximumPopulation(logs)
	if res != exp {
		t.Fail()
	}
}

func TestMaximumPopulation2(t *testing.T) {
	logs := [][]int{{1950, 1961}, {1960, 1971},{1970,1981}}
	exp := 1960
	res := MaximumPopulation(logs)
	if res != exp {
		t.Fail()
	}
}
