package solution

import (
	"math"
	"testing"
)

func TestTrimMean(t *testing.T) {
	arr := []int{6, 2, 7, 5, 1, 2, 0, 3, 10, 2, 5, 0, 5, 5, 0, 8, 7, 6, 8, 0}
	exp := 4.0
	if math.Abs(TrimMean(arr)-exp) > 0.0001 {
		t.Fail()
	}
}

func TestTrimMean2(t *testing.T) {
	arr := []int{6, 0, 7, 0, 7, 5, 7, 8, 3, 4, 0, 7, 8, 1, 6, 8, 1, 1, 2, 4, 8, 1, 9, 5, 4, 3, 8, 5, 10, 8, 6, 6, 1, 0, 6, 10, 8, 2, 3, 4}
	exp := 4.77778
	if math.Abs(TrimMean(arr)-exp) > 0.0001 {
		t.Fail()
	}
}

func TestTrimMean3(t *testing.T) {
	arr := []int{9, 7, 8, 7, 7, 8, 4, 4, 6, 8, 8, 7, 6, 8, 8, 9, 2, 6, 0, 0, 1, 10, 8, 6, 3, 3, 5, 1, 10, 9, 0, 7, 10, 0, 10, 4, 1, 10, 6, 9, 3, 6, 0, 0, 2, 7, 0, 6, 7, 2, 9, 7, 7, 3, 0, 1, 6, 1, 10, 3}
	exp := 5.27778
	if math.Abs(TrimMean(arr)-exp) > 0.0001 {
		t.Fail()
	}
}
