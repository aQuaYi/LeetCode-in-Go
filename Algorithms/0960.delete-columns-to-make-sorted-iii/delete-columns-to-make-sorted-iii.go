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
		for _, r := range res {
			if s[r[len(r)-1]] <= s[i] {
				r = append(r, i)
				isAdded = true
			}
		}
		if !isAdded {
			t := make([]int, 1, size)
			t[0] = i
			res = append(res, t)
		}
	}

	return res

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
