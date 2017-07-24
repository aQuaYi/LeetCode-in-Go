package Problem0010

import (
	"strings"
)

// 程序中存在以下假设
// "*" 不会出现在p的首位
// "**" 不会出现，但会出现 ".*."" , ".*.." , ".*.*"

// 从p的后面，往前检查。

func isMatch(s string, p string) bool {
	if len(s) == 0 && len(p) == 0 {
		return true
	}

	if s[len(s)-1] == p[len(p)-1] || p[len(p)-1] == '.' {
		return isMatch(s[:len(s)-1], p[:len(p)-1])
	}

	if p[len(p)-1] == '*' { // 此时，len(p) >= 2
		switch {
		case p[len(p)-2] == '.':
			if len(p) == 2 {
				return true
			}
			switch { // 此时， len(p) >= 3
			case p[len(p)-3] == '.':
			switch{
				case p[len(p)-4]==
			}
			}
		default:
			if s[len(s)-1] != p[len(p)-2] {
				return false
			}
			if strings.Contains(s, p[len(p)-2:len(p)-1]) {
				return isMatch(strEndBefore(s, p[len(p)-2]), p[:len(p)-2])
			}
			return false
		}

	}

	return true
}

func strEndBefore(s string, b byte) string {

	for i := len(s) - 1; i > 0; i-- {
		if s[i] == b && s[i-1] != b {
			return s[:i]
		}
	}

	return ""
}
