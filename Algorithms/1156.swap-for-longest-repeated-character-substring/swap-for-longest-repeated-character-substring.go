package problem1156

import "strings"

func maxRepOpt1(text string) int {
	rec := [26][][]int{}
	segments, index := split(text), 0
	for _, s := range segments {
		b, n := int(s[0]-'a'), len(s)
		rec[b] = append(rec[b], []int{index, index + n + 1, n})
		index += n
	}

	res := 0
	for _, r := range rec {
		n := len(r)
		if n == 0 {
			continue
		}
		ext := min(n-1, 1)         // extension a lonely segment
		con := max(0, min(n-2, 1)) // connect two neighbor segments
		prev := r[0]
		res = max(res, prev[2]+ext)
		for i := 1; i < len(r); i++ {
			cur := r[i]
			if prev[1] == cur[0] {
				res = max(res, prev[2]+cur[2]+con)
			} else {
				res = max(res, cur[2]+ext)
			}
			prev = cur
		}
	}

	return res
}

func split(s string) []string {
	var sb strings.Builder
	p := s[0]
	sb.WriteByte(p)
	for i := 1; i < len(s); i++ {
		n := s[i]
		if p != n {
			sb.WriteByte('\n')
		}
		sb.WriteByte(n)
		p = n
	}
	return strings.Split(sb.String(), "\n")
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
