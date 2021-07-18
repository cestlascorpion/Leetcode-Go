package solution

import "testing"

func TestMaximumTime(t *testing.T) {
	time := []string{"2?:?0", "0?:3?", "1?:22", "?4:03"}
	exp := []string{"23:50", "09:39", "19:22", "14:03"}
	for i := range time {
		if MaximumTime(time[i]) != exp[i] {
			t.Fail()
		}
	}
}
