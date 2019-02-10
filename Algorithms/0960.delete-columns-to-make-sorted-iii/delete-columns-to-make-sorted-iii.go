package problem0960

import "sort"

func minDeletionSize(A []string) int {
	size, wordSize := len(A), len(A[0])
	keeps := getKeeps(A[0])

	sort.Slice(keeps, func(i int, j int) bool {
		return len(keeps[i]) > len(keeps[j])
	})

	for _, k := range keeps {
		isAllLexicographic := true
		for i := 1; i < size; i++ {
			if !isLexicographic(A[i], k) {
				isAllLexicographic = false
				break
			}
		}
		if isAllLexicographic {
			return wordSize - len(k)
		}
	}
	return wordSize - 1
}

func getKeeps(s string) [][]int {
	size := len(s)
	res := make([][]int, 1, size)
	res[0] = make([]int, 1, size)

	for i := 1; i < size; i++ {
		isAdded := false
		for j := range res {
			if s[res[j][len(res[j])-1]] <= s[i] {
				res[j] = append(res[j], i)
				isAdded = true
			}
		}
		if !isAdded {
			t := make([]int, 1, size)
			t[0] = i
			res = append(res, t)
		}
	}

	t := make([][]int, 0, size)
	for i := 0; i < len(res); i++ {
		for j := i + 1; j < len(res); j++ {
			set := intersection(res[i], res[j])
			if len(set) > 0 {
				t = append(t, set)
			}
		}
	}

	return append(res, t...)

}

func isLexicographic(s string, keep []int) bool {
	size := len(keep)
	a := s[keep[0]]
	for i := 1; i < size; i++ {
		b := s[keep[i]]
		if a > b {
			return false
		}
		a = b
	}
	return true
}

func intersection(a, b []int) []int {
	res := make([]int, 0, 100)
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		m, n := a[i], b[j]
		switch {
		case m < n:
			i++
		case m > n:
			j++
		default:
			res = append(res, m)
			i++
			j++
		}
	}
	return res
}
