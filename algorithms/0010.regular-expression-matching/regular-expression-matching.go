package Problem0010

import (
	"strings"
)

// 程序中存在以下假设
// "*" 不会出现在p的首位
// ".**" 不会出现，但会出现 ".*."" , ".*.." , ".*.*"

// 从p的后面，往前检查。

func isMatch(s string, p string) bool {
	if len(s) == 0 && len(p) == 0 {
		return true
	}

	switch {
	case p[0:2] == ".*":
		return dealDotStar(s, p)
	case s[0] == p[0]:
		return isMatch(s[1:], p[1:])
	case s[0] != p[0] && p[0] == p[0]:

	}

	if strings.Contains(p, ".*") {
		return dealDotStar(s, p)
	}

	return true
}

func dealDotStar(s, p string) bool {
	switch {
	case len(p) == 2:
		return true

	}

	if len(p) == 2 {
		return true
	}

	// i 是p中第一个".*"的后一个字符的位置
	i := 2
	for i < len(p) {
		if p[i-2:i] == ".*" {
			break
		}
	}
	if i == len(p) {
		return true
	}

	remain := len(p) - i

	return isMatch(s[len(s)-remain:], p[i:])
}

func stringBegin(s string, c byte) string {

	for i := 0; i < len(s); i++ {
		if s[i] == c {
			return s[i:]
		}
	}

	return ""
}

// for i := 0; i < len(p); i++ {

// 	for j := 0; j < len(s); j++ {
// 		switch {
// 		case p[i] == s[j] || p[i] == '.':
// 			continue
// 		case p[i] == '*' && (p[i-1] == '.' || p[i-1] == s[j]):
// 			continue

// 		}
// 	}
// }

// 	i, j := 0, 0
// 	for {
// 		switch {
// 		case s[i] == p[j] || p[j] == '.':
// 			i++
// 			j++
// 		case p[j] == '*':
// 			if p[j-1] == '.' || p[j-1] == s[i] {

// 			}
// 		}

// 		if i < len(s) && j == len(p) {
// 			return false
// 		}

// 		if i == len(s) {
// 			return true
// 		}
// 	}

// 	panic("NEVER BE HERE")
// }

// func dealStar(s, p string, i, j int) (int, int) {

// }
