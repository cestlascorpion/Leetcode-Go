package solution

const (
	WinS1 = 448 // 111000000
	WinS2 = 56  // 000111000
	WinS3 = 7   // 000000111
	WinS4 = 292 // 100100100
	WinS5 = 146 // 010010010
	WinS6 = 73  // 001001001
	WinS7 = 273 // 100010001
	WinS8 = 84  // 001010100
)

func win(now uint) bool {
	return (now&WinS1) == WinS1 || (now&WinS2) == WinS2 || (now&WinS3) == WinS3 || (now&WinS4) == WinS4 || (now&WinS5) == WinS5 || (now&WinS6) == WinS6 || (now&WinS7) == WinS7 || (now&WinS8) == WinS8
}

func now(shape uint, i, j int) uint {
	return shape | (1 << (9 - i*3 - j - 1))
}

func Tictactoe(moves [][]int) string {
	data := [2]uint{0, 0}
	for i := range moves {
		num := now(data[i%2], moves[i][0], moves[i][1])
		if win(num) {
			if i%2 == 0 {
				return "A"
			} else {
				return "B"
			}
		}
		data[i%2] = num
	}
	if len(moves) == 9 {
		return "Draw"
	} else {
		return "Pending"
	}
}
