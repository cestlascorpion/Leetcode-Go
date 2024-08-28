package solution

import (
	"strings"
	"testing"
)

func splitEmail(email string) (string, string) {
	var local, domain string
	for i := range email {
		if email[i] == '@' {
			local = email[:i]
			domain = email[i:]
			break
		}
	}
	local = strings.ReplaceAll(local, ".", "")
	for i := 0; i < len(local); i++ {
		if local[i] == '+' {
			local = local[:i]
			break
		}
	}
	return local, domain
}

func NumUniqueEmails(emails []string) int {
	m := make(map[string]int)
	for _, email := range emails {
		local, domain := splitEmail(email)
		m[local+domain]++
	}
	return len(m)
}

func TestNumUniqueEmails(t *testing.T) {
	emails := []string{"a@leetcode.com", "b@leetcode.com", "c@leetcode.com"}
	exp := 3
	if NumUniqueEmails(emails) != exp {
		t.Fail()
	}

	emails = []string{"test.emailalex@leetcode.com", "test.e.mailbob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}
	exp = 3
	if NumUniqueEmails(emails) != exp {
		t.Fail()
	}
}
