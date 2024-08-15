package solution

import "testing"

func IsSubsequence(s string, t string) bool {
	i, j := 0, 0 // double pointer
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == len(s)
}

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want bool
	}{
		{"abc", "ahbgdc", true},
		{"axc", "ahbgdc", false},
	}
	for _, tt := range tests {
		if got := IsSubsequence(tt.s, tt.t); got != tt.want {
			t.Errorf("IsSubsequence(%v, %v) = %v, want %v", tt.s, tt.t, got, tt.want)
		}
	}
}