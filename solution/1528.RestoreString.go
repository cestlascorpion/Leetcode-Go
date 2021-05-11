package solution

func RestoreString(s string, indices []int) string {
	bytes := make([]byte, len(s))
	for i := range indices {
		bytes[indices[i]] = s[i]
	}
	return string(bytes)
}
