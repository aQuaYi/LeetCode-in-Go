package Problem0331

import (
	"strings"
)

func isValidSerialization(preorder string) bool {
	if preorder == "#" {
		return true
	}

	ss := strings.Split(preorder, ",")

	if (len(ss)-3)%2 != 0 {
		return false
	}

	return check(ss)
}

func check(ss []string) bool {
	n := len(ss)
	if n < 3 || ss[0] == "#" {
		return false
	}

	if n == 3 &&
		ss[1] == "#" &&
		ss[2] == "#" {
		return true
	}

	if ss[0] != "#" &&
		ss[1] == "#" &&
		ss[2] != "#" {
		return check(ss[2:])
	}

	if ss[n-3] == "#" &&
		ss[n-2] == "#" &&
		ss[n-1] == "#" {
		return check(ss[1 : n-1])
	}

	ss[n-3] = "#"
	return check(ss[:n-2])
}
