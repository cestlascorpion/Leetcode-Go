package solution

type pathPoint struct {
	x int
	y int
}

func IsPathCrossing(path string) bool {
	x, y := 0, 0
	points := make(map[pathPoint]struct{})
	points[pathPoint{x, y}] = struct{}{}
	for i := range path {
		switch path[i] {
		case 'N':
			y++
			p := pathPoint{x, y}
			if _, ok := points[p]; !ok {
				points[p] = struct{}{}
			} else {
				return true
			}
		case 'S':
			y--
			p := pathPoint{x, y}
			if _, ok := points[p]; !ok {
				points[p] = struct{}{}
			} else {
				return true
			}
		case 'E':
			x++
			p := pathPoint{x, y}
			if _, ok := points[p]; !ok {
				points[p] = struct{}{}
			} else {
				return true
			}
		case 'W':
			x--
			p := pathPoint{x, y}
			if _, ok := points[p]; !ok {
				points[p] = struct{}{}
			} else {
				return true
			}
		}
	}
	return false
}
