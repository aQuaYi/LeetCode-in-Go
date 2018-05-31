package problem0833

import (
	"bytes"
)

func findReplaceString(S string, indexes []int, sources []string, targets []string) string {
	size := len(S)

	ss := make([]string, 0, len(indexes)*2+2)

	end := 0
	for i, idx := range indexes {
		begin := end
		ss = append(ss, S[begin:idx])
		end = min(size, idx+len(sources[i]))
		ss = append(ss, S[idx:end])
	}

	if end < size {
		ss = append(ss, S[end:])
	}

	var buf bytes.Buffer

	for i, s := range ss {
		if i%2 == 0 {
			buf.WriteString(s)
			continue
		}

		if s == sources[i/2] {
			buf.WriteString(targets[i/2])
		} else {
			buf.WriteString(s)
		}
	}

	return buf.String()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
