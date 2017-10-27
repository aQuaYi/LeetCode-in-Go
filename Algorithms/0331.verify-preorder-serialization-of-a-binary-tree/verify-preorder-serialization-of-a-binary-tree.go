package Problem0331

import (
	"strings"
)

func isValidSerialization(preorder string) bool {
	if preorder == "#" {
		return true
	}

	preorder = strings.Replace(preorder, ",", "", -1)

	if (len(preorder)-3)%2 != 0 {
		return false
	}

	return check(preorder)
}

func check(bs string) bool {
	n := len(bs)
	if n < 3 || bs[0] == '#' {
		return false
	}

	if n == 3 && bs[1:] == "##" {
		return true
	}

	if bs[n-3:] == "###" {
		return check(bs[1 : n-1])
	}

	return check(bs[:n-3] + "#")
}
