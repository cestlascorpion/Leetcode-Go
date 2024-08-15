package solution

import "testing"

func LemonadeChange(bills []int) bool {
	five, ten := 0, 0 // only need to count 5 and 10
	for _, bill := range bills {
		if bill == 5 {
			five++
		} else if bill == 10 {
			if five == 0 {
				return false
			}
			five--
			ten++
		} else { // 20
			if ten > 0 && five > 0 {
				ten--
				five--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}

func TestLemonadeChange(t *testing.T) {
	tests := []struct {
		bills []int
		want  bool
	}{
		{[]int{5, 5, 5, 10, 20}, true},
		{[]int{5, 5, 10}, true},
		{[]int{10, 10}, false},
		{[]int{5, 5, 10, 10, 20}, false},
	}
	for _, tt := range tests {
		if got := LemonadeChange(tt.bills); got != tt.want {
			t.Errorf("LemonadeChange(%v) = %v, want %v", tt.bills, got, tt.want)
		}
	}
}