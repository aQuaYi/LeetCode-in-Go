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

func clean(e string) string {
	username, atDomain := split(e)
	return cancelDot(cutTail(username)) + atDomain
}

func split(e string) (username, atDomain string) {
	i := strings.IndexByte(e, '@')
	return e[:i], e[i:]
}

func cutTail(username string) string {
	i := strings.IndexByte(username, '+')
	if i == -1 {
		return username
	}
	return username[:i]
}

func cancelDot(username string) string {
	return strings.Replace(username, ".", "", -1)
}
