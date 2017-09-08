package Problem0093

import (
	"fmt"
)

func restoreIpAddresses(s string) []string {
	n := len(s)
	if n < 4 || n > 12 {
		return []string{}
	}

	res := []string{}
	combination := []string{}
	var dfs func(int, int)

	dfs = func(idx, begin int) {
		if idx == 3 {
			combination[3] = s[begin:]
			if isIP(combination) {
				res = append(res, IP(combination))
			}
			return
		}

		for i := begin + 1; i <= n-(3-idx); i++ {
			if n-i > 3*(3-idx) {
				// 后面的 IP 段 至少有一个超过了 3 个字符
				// 说明此IP段短了
				continue
			}

			if i-begin > 3 {
				// 此 IP 段长度，超过 3 位了
				break
			}

			combination[idx] = s[begin:i]
			dfs(idx+1, i)
		}

	}

	dfs(0, 0)

	return res
}

// IP 返回 ss 代表的 IP 地址
func IP(ss []string) string {
	return fmt.Sprintf("%s.%s.%s.%s", ss[0], ss[1], ss[2], ss[3])
}

func isIP(ss []string) bool {
	for i := range ss {
		if !isOK(ss[i]) {
			return false
		}
	}

	return true
}

func isOK(s string) bool {
	if len(s) > 1 && s[0] == '0' || len(s) > 3 {
		return false
	}

	if len(s) < 3 {
		return true
	}

	switch s[0] {
	case '1':
		return true
	case '2':
		if '0' <= s[1] && s[1] <= '4' {
			return true
		}
		if s[1] == '5' && '0' <= s[2] && s[2] <= '5' {
			return true
		}
		return false
	default:
		return false
	}
}
