package solution

func CountGoodRectangles(rectangles [][]int) int {
	m := make(map[int]int)
	for i := range rectangles {
		l := rectangles[i][0]
		w := rectangles[i][1]
		if l > w {
			m[w]++
		} else {
			m[l]++
		}
	}
	max := 0
	sum := 0
	for k, v := range m {
		if k > max {
			max = k
			sum = v
		}
	}
	return sum
}
