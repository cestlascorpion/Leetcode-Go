package solution

func CountStudents(students []int, sandwiches []int) int {
	sad := 0
	deadlock := true
	for {
		deadlock = true
		for i := 0; i < len(students); i++ {
			if students[i] == -1 {
				continue
			}
			if students[i] == sandwiches[sad] {
				students[i] = -1
				deadlock = false
				sad++
			}
		}
		if deadlock {
			break
		}
	}
	return len(sandwiches) - sad
}
