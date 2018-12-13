package problem0929

import (
	"strings"
)

func numUniqueEmails(emails []string) int {
	hasSeen := make(map[string]bool, 100)
	for _, email := range emails {
		hasSeen[clean(email)] = true
	}
	return len(hasSeen)
}

func clean(email string) string {
	user, atDomain := split(email)
	return cancelDot(trim(user)) + atDomain
}

func split(email string) (user, atDomain string) {
	i := strings.IndexByte(email, '@')
	return email[:i], email[i:]
}

func cancelDot(user string) string {
	return strings.Replace(user, ".", "", -1)
}

func trim(user string) string {
	i := strings.IndexByte(user, '+')
	if i == -1 {
		return user
	}
	return user[:i]
}
