package Problem0076

func minWindow(s string, t string) string {
	rec := make(map[byte][]int, len(s))

	for i := range s {
		rec[s[i]] = append(rec[s[i]], i)
	}

	// 检查，t 中的字母是否都在 s 中
	for i := range t {
		if len(rec[t[i]]) == 0 {
			return ""
		}
	}

	return ""
}
