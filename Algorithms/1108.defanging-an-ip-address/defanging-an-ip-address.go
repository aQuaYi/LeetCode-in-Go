package problem1108

import (
	"strings"
)

func defangIPaddr(address string) string {
	return strings.Replace(address, ".", "[.]", 4)
}

func defangIPaddr2(address string) string {
	var res string
	for _, v := range address {
		if v == '.' {
			res += "[.]"
		} else {
			res += string(v)
		}
	}
	return res
}
