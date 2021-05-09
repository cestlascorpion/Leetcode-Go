package solution

func LargestAltitude(gain []int) int {
	altitude := 0
	cursor := 0
	for _, delta := range gain {
		cursor += delta
		if cursor > altitude {
			altitude = cursor
		}
	}
	return altitude
}
