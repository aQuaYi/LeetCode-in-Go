package problem0929

import (
	"strings"
)

func numUniqueEmails(emails []string) int {
	count := make(map[string]bool, 100)
	for _, e := range emails {
		count[clean(e)] = true
	}
	return len(count)
}

func clean(email string) string {
	i := strings.IndexByte(email, '@')
	return deleteDot(trim(email[:i])) + email[i:]
}

func deleteDot(username string) string {
	return strings.Replace(username, ".", "", -1)
}

func trim(username string) string {
	i := strings.IndexByte(username, '+')
	if i == -1 {
		return username
	}
	return username[:i]
}
