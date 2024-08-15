package solution

import (
	"math"
	"testing"
)

func ArrangeCoins(n int) int {
	// 1 + 2 + 3 + ... + k = k(k+1)/2 通项公式
	// k^2 + k = 2n -> sqrt(2n) > k > sqrt(2n) - 1
	i := int(math.Round(math.Sqrt(float64(2 * n))))
	if i*(i+1) <= 2*n {
		return i
	}
	return i - 1
}

func TestArrangeCoins(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 1},
		{5, 2},
		{8, 3},
	}
	for _, tt := range tests {
		if got := ArrangeCoins(tt.n); got != tt.want {
			t.Errorf("ArrangeCoins(%v) = %v, want %v", tt.n, got, tt.want)
		}
	}
}
