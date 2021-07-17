package solution

func SlowestKey(releaseTimes []int, keysPressed string) byte {
	dict := [26]int{}
	for i := range dict {
		dict[i] = -1
	}
	prev := 0
	for i := range releaseTimes {
		idx := keysPressed[i] - 'a'
		if dict[idx] < releaseTimes[i]-prev {
			dict[idx] = releaseTimes[i] - prev
		}
		prev = releaseTimes[i]
	}
	res := keysPressed[0]
	max := releaseTimes[0]
	for i := range dict {
		if dict[i] > max {
			max = dict[i]
			res = uint8(i) + 'a'
		} else if dict[i] == max {
			if uint8(i)+'a' > res {
				res = uint8(i) + 'a'
			}
		}
	}
	return res
}
