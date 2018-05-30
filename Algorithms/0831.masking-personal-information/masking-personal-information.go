package problem0831

import "strings"

func maskPII(s string) string {
	if strings.ContainsRune(s, '@') {
		return maskEmailAddress(strings.ToLower(s))
	}
	return maskPhoneNumber(s)
}

func maskEmailAddress(s string) string {
	ss := strings.Split(s, "@")
	ss[0] = ss[0][0:1] + "*****" + ss[0][len(ss[0])-1:]
	return strings.Join(ss, "@")
}

func maskPhoneNumber(s string) string {
	tmp := make([]byte, 0, 13)
	for i := range s {
		if '0' <= s[i] && s[i] <= '9' {
			tmp = append(tmp, s[i])
		}
	}

	size := len(tmp)
	s = string(tmp[size-4:])

	if size == 10 {
		return "***-***-" + s
	}

	return "+" + strings.Repeat("*", size-10) + "-***-***-" + s
}
