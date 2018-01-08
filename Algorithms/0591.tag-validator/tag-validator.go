package Problem0591

import (
	"strings"
)

func isValid(code string) bool {
	stack := make([]string, 0, 16)
	for i := 0; i < len(code); {
		if i > 0 && len(stack) == 0 {
			return false
		}

		if strings.HasPrefix(code[i:], "<![CDATA[") {
			j := i + 9
			i = strings.Index(code[j:], "]]>") + j
			if i < j {
				return false
			}
			i += 3
		} else if strings.HasPrefix(code[i:], "</") {
			j := i + 2
			i = strings.Index(code[j:], ">") + j
			if i < j || // 没有 ">"
				i == j || // TAG_NAME 为 ""
				i-j > 9 { // TAG_NAME 长度超过 9
				return false
			}
			for k := j; k < i; k++ {
				if code[k] < 'A' || 'Z' < code[k] {
					return false
				}
			}
			tagName := code[j:i]
			i++
			n := len(stack)
			if n == 0 || stack[n-1] != tagName {
				return false
			}
			stack = stack[:n-1]
		} else if code[i] == '<' {
			j := i + 1
			i := j + strings.Index(code[j:], ">")
			if i < j || // 没有 ">"
				i == j || // TAG_NAME 为 ""
				i-j > 9 { // TAG_NAME 长度超过 9
				return false
			}
			for k := j; k < i; k++ {
				if code[k] < 'A' || 'Z' < code[k] {
					return false
				}
			}
			tagName := code[j:i]
			i++
			stack = append(stack, tagName)
		} else {
			i++
		}

	}

	return len(stack) == 0
}
