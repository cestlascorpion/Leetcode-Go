package solution

func BusyStudent(startTime []int, endTime []int, queryTime int) int {
	if len(startTime) != len(endTime) {
		panic("start time size != end time size")
	}
	busy := 0
	n := len(startTime)
	for i := 0; i < n; i++ {
		if queryTime >= startTime[i] && queryTime <= endTime[i] {
			busy++
		}
	}
	return busy
}
