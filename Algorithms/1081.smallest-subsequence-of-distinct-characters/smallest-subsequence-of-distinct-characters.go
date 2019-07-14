package problem1081

import (
	"sort"
	"strings"
)

func smallestSubsequence(text string) string {
	rec := make([][]int, 26)
	for i, b := range text {
		c := b - 'a'
		rec[c] = append(rec[c], i)
	}

	letters := make([]*letter, 0, 26)
	for i := 0; i < 26; i++ {
		if len(rec[i]) == 0 {
			continue
		}
		letters = append(letters, &letter{
			char:   byte(i + 'a'),
			indexs: rec[i],
		})
	}

	sort.Slice(letters, func(i int, j int) bool {
		return isLess(letters[i], letters[j])
	})

	var sb strings.Builder
	for _, l := range letters {
		sb.WriteByte(l.char)
	}

	return sb.String()
}

type letter struct {
	char   byte
	indexs []int
}

func isLess(a, b *letter) bool {
	return a.char < b.char && a.indexs[0] < b.indexs[len(b.indexs)-1]
}
