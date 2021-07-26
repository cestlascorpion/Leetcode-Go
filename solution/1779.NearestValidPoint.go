package solution

import "math"

func NearestValidPoint(x int, y int, points [][]int) int {
	found := false
	index := -1
	delta := math.MaxInt32
	for idx := range points {
		if points[idx][0] == x || points[idx][1] == y {
			lx := points[idx][0] - x
			if lx < 0 {
				lx = -lx
			}
			ly := points[idx][1] - y
			if ly < 0 {
				ly = -ly
			}
			if !found {
				found = true
				index = idx
				delta = lx + ly
			} else {
				if lx+ly < delta {
					index = idx
					delta = lx + ly
				}
			}
		}
	}
	return index
}
