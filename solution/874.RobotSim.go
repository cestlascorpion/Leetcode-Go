package solution

const (
	KUp           = 0
	KLeft         = 1
	KDown         = 2
	KRight        = 3
	KDirectionMax = 4
)

type point struct {
	x int
	y int
}

func turn(direction, command int) int {
	switch command {
	case -2:
		return (direction + 1) % KDirectionMax
	case -1:
		return (direction + KDirectionMax - 1) % KDirectionMax
	}
	return direction
}

func path(x, y, direction, length int, obstacles map[point]struct{}) (int, int) {
	switch direction {
	case KUp:
		for i := 1; i <= length; i++ {
			p := point{x, y + 1}
			if _, ok := obstacles[p]; ok {
				break
			}
			y++
		}
	case KLeft:
		for i := 1; i <= length; i++ {
			p := point{x - 1, y}
			if _, ok := obstacles[p]; ok {
				break
			}
			x--
		}
	case KDown:
		for i := 1; i <= length; i++ {
			p := point{x, y - 1}
			if _, ok := obstacles[p]; ok {
				break
			}
			y--
		}
	case KRight:
		for i := 1; i <= length; i++ {
			p := point{x + 1, y}
			if _, ok := obstacles[p]; ok {
				break
			}
			x++
		}
	}
	return x, y
}

func RobotSim(commands []int, obstacles [][]int) int {
	res := 0
	dict := make(map[point]struct{})
	for i := range obstacles {
		dict[point{obstacles[i][0], obstacles[i][1]}] = struct{}{}
	}
	x, y := 0, 0
	dir := KUp
	for i := range commands {
		if commands[i] < 0 {
			dir = turn(dir, commands[i])
		} else {
			x, y = path(x, y, dir, commands[i], dict)
		}
		if distance := x*x + y*y; distance > res {
			res = distance
		}
	}
	return res
}
