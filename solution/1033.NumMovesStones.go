package solution

import (
	"sort"
)

func maxMoves(a int, b int, c int) int {
	array := []int{a, b, c}
	sort.Ints(array)
	return array[2] - array[0] - 2
}

func minMoves(a int, b int, c int) int {
	array := []int{a, b, c}
	sort.Ints(array)
	x, y := array[1]-array[0], array[2]-array[1]
	if x == 1 && y == 1 {
		return 0
	}
	if x == 1 || y == 1 {
		return 1
	}
	if x == 2 || y == 2 {
		return 1
	}
	return 2
}

func NumMovesStones(a int, b int, c int) []int {
	return []int{minMoves(a, b, c), maxMoves(a, b, c)}
}
