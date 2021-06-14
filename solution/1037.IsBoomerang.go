package solution

type vector struct {
	x int
	y int
}

func IsBoomerang(points [][]int) bool {
	v1 := vector{
		x: points[1][0] - points[0][0],
		y: points[1][1] - points[0][1],
	}
	v2 := vector{
		x: points[2][0] - points[0][0],
		y: points[2][1] - points[0][1],
	}
	return v1.x*v2.y != v1.y*v2.x
}
