package solution

import "testing"

func TestRunningSum(t *testing.T) {
	numss := [][]int{{1, 2, 3, 4}, {1, 1, 1, 1, 1}}
	expss := [][]int{{1, 3, 6, 10}, {1, 2, 3, 4, 5}}
	for i := range numss {
		sums := RunningSum(numss[i])
		if len(expss[i]) != len(sums) {
			t.Fatalf("exp: %v sums: %v\n", expss[i], sums)
		}
		for j := range expss[i] {
			if expss[i][j] != sums[j] {
				t.Errorf("exp: %d sum: %d\n", expss[i][j], sums[j])
			}
		}
	}

}
