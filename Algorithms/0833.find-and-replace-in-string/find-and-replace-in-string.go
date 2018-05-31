package problem0833

import (
	"bytes"
	"sort"
)

type entry struct {
	idx            int
	source, target string
}

func findReplaceString(S string, indexes []int, sources []string, targets []string) string {
	entries := make([]*entry, len(indexes))
	for i := range entries {
		entries[i] = &entry{
			idx:    indexes[i],
			source: sources[i],
			target: targets[i],
		}
	}

	sort.Slice(entries, func(i int, j int) bool {
		return entries[i].idx < entries[j].idx
	})

	size := len(S)

	ss := make([]string, 0, len(indexes)*2+2)

	end := 0
	for i := range entries {
		begin := end
		ss = append(ss, S[begin:entries[i].idx])
		end = min(size, entries[i].idx+len(entries[i].source))
		ss = append(ss, S[entries[i].idx:end])
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

		if s == entries[i/2].source {
			buf.WriteString(entries[i/2].target)
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
