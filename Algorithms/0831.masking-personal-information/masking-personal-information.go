package problem0831

import "strings"

func maskPII(s string) string {
	if strings.ContainsRune(s, '@') {
		return maskEmailAddress(strings.ToLower(s))
	}
	return maskPhoneNumber(parse(s))
}

func maskEmailAddress(s string) string {
	ss := strings.Split(s, "@")
	ss[0] = ss[0][0:1] + "*****" + ss[0][len(ss[0])-1:]
	return strings.Join(ss, "@")
}

func maskPhoneNumber(size int, s string) string {
	if size == 10 {
		return "***-***-" + s
	}
	return "+" + strings.Repeat("*", size-10) + "-***-***-" + s
}

// 返回 s 中数字的个数，和，s 中最后 4 个数字
func parse(s string) (int, string) {
	tmp := make([]byte, 0, 13)
	for i := range s {
		if '0' <= s[i] && s[i] <= '9' {
			tmp = append(tmp, s[i])
		}
	}
	size := len(tmp)
	return size, string(tmp[size-4:])
}
