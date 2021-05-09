package solution

func Calculate(s string) int {
	x := 1
	y := 0
	for i := range s {
		switch s[i] {
		case 'A':
			x = 2*x + y
		case 'B':
			y = 2*y + x
		default:
			panic("invalid char")
		}
	}
	return x + y
}
