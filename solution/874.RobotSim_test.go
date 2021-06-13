package solution

import "testing"

func TestRobotSim(t *testing.T) {
	commands := []int{4, -1, 4, -2, 4}
	obstacles := [][]int{{2, 4}}
	exp := 65
	if RobotSim(commands, obstacles) != exp {
		t.Fail()
	}
}

func TestRobotSim2(t *testing.T) {
	commands := []int{4, -1, 3}
	obstacles := [][]int{}
	exp := 25
	if RobotSim(commands, obstacles) != exp {
		t.Fail()
	}
}

func TestRobotSim3(t *testing.T) {
	commands := []int{-2, -1, 8, 9, 6}
	obstacles := [][]int{{-1, 3}, {0, 1}, {-1, 5}, {-2, -4}, {5, 4}, {-2, -3}, {5, -1}, {1, -1}, {5, 5}, {5, 2}}
	exp := 0
	if RobotSim(commands, obstacles) != exp {
		t.Fail()
	}
}
