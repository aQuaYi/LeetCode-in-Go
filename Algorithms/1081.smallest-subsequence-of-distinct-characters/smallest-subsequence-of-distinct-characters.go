package problem1081

import "strings"

func smallestSubsequence(text string) string {
	rec := make([][]int, 26)
	count := 0
	for i, b := range text {
		c := b - 'a'
		rec[c] = append(rec[c], i)
		if len(rec[c]) == 1 {
			count++
		}
	}

	flag := -1
	beforeAll := func() bool {
		ok := true
		for i := 0; i < 26 && ok; i++ {
			if len(rec[i]) == 0 {
				continue
			}

		}
		return ok
	}

	var sb strings.Builder
	for i := 0; i < count; i++ {
		for j := 0; j < 26; j++ {
			if len(rec[i]) == 0 {
				continue
			}
			j := 0
			for j < len(rec[i]) && rec[i][j] < flag {
				j++
			}
			index := rec[i][0]
		}

	}

	return sb.String()
}
