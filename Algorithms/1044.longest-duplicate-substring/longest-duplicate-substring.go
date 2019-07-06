package problem1044

func longestDupSubstring(S string) string {
	n := len(S)
	chars := make([][]int, 26)
	for i, r := range S {
		b := int(r - 'a')
		chars[b] = append(chars[b], i)
	}
	res := ""
	alpha := numberify(S)
	var dfs func(string, int, []int)
	dfs = func(s string, b int, indexs []int) {
		incIndexs := inc(indexs)
		seen := [26]bool{}
		for _, i := range indexs {
			if i+1 >= n {
				continue
			}
			c := alpha[i+1]
			if seen[c] {
				continue
			}
			next := intersect(incIndexs, chars[c])
			if len(next) < 2 {
				continue
			}
			t := s + string(c+'a')
			if len(res) <= len(s) {
				res = t
			}
			dfs(t, c, next)
			seen[c] = true
		}
	}

	us := unique(S)
	for _, b := range us {
		dfs(string(b+'a'), b, chars[b])
	}

	return res
}

func unique(S string) []int {
	res := make([]int, 0, len(S))
	has := [26]bool{}
	for _, r := range S {
		b := int(r - 'a')
		if has[b] {
			continue
		}
		res = append(res, b)
		has[b] = true
	}
	return res
}

func inc(A []int) []int {
	res := make([]int, len(A))
	for i := range A {
		res[i] = A[i] + 1
	}
	return res
}

func intersect(A, B []int) []int {
	m, n := len(A), len(B)
	res := make([]int, 0, min(m, n))
	i, j := 0, 0
	for i < m && j < n {
		switch {
		case A[i] < B[j]:
			i++
		case A[i] > B[j]:
			j++
		default:
			res = append(res, A[i])
			i++
			j++
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func numberify(S string) []int {
	res := make([]int, len(S))
	for i, r := range S {
		res[i] = int(r - 'a')
	}
	return res
}
