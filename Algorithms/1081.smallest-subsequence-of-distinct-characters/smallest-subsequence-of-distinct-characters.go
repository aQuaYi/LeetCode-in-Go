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

	clean := func(flag int) {
		for i := 0; i < 26; i++ {
			if len(rec[i]) == 0 {
				continue
			}
			j := 0
			for j < len(rec[i]) && rec[i][j] < flag {
				j++
			}
			rec[i] = rec[i][j:]
		}
	}

	beforeAll := func(index int) bool {
		ok := true
		for i := 0; i < 26 && ok; i++ {
			if len(rec[i]) == 0 {
				continue
			}
			ok = index <= rec[i][len(rec[i])-1]
		}
		return ok
	}

	var sb strings.Builder
	for i := 0; i < count; i++ {
		for j := 0; j < 26; j++ {
			if len(rec[j]) == 0 {
				continue
			}
			index := rec[j][0]
			if beforeAll(index) {
				sb.WriteByte(byte(j + 'a'))
				rec[j] = nil
				clean(index)
				break
			}
		}
	}

	return sb.String()
}
