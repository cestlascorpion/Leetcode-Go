package solution

import "time"

const KGuessParamN = 10

var GuessNum int

func init() {
	GuessNum = int(time.Now().Unix())%KGuessParamN + 1
}

func guess(num int) int {
	if num == GuessNum {
		return 0
	} else if num < GuessNum {
		return 1
	} else {
		return -1
	}
}

func GetGuessNum() int {
	return GuessNum
}

func GuessNumber(n int) int {
	l, r := 1, n
	for l < r {
		m := l + (r-l)/2
		switch guess(m) {
		case 0:
			return m
		case -1:
			r = m
		case 1:
			l = m + 1
		}
	}
	return l
}
