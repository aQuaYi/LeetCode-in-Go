package problem1044

func longestDupSubstring(S string) string {
	n := len(S)
	chars := make([][]int, 26)
	for i, r := range S {
		b := int(r - 'a')
		chars[b] = append(chars[b], i)
	}

	alpha := numberify(S)
	var dfs func(int)
	dfs = func(b int) {
		indexs := chars[b]
		incIndexs := inc(indexs)
		seen := [26]bool{}
		for k := 0; k < len(indexs); k++ {
			i := indexs[k]
			c := alpha[i+1]
			if seen[c] {
				continue
			}
			next := intersect(incIndexs, chars[c])

			seen[c] = true
		}
	}
	res := ""
	for c := 0; c < 26; c++ {
		indexs := chars[c]
		for i := 0; i < len(indexs); i++ {
			if indexs[i]+1 == n {
				continue
			}
			next := S[indexs[i]+1]
			tmp := make([]int, 0, len(indexs))
			for j := 0; j < len(indexs); j++ {
				if indexs[j]+1 < n && next == S[indexs[j]+1] {
					tmp = append(tmp, indexs[j]+1)
				}
			}

		}
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
	for i := range A {
		A[i]++
	}
	return A
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
