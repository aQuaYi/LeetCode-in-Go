package Problem0767

import (
	"sort"
)

func reorganizeString(s string) string {
	n := len(s)
	bs := []byte(s)

	count := make([][2]int, 26)
	maxCount := 0
	for _, b := range bs {
		count[b-'a'][0]++
		count[b-'a'][1] = int(b - 'a')
		maxCount = max(maxCount, count[b-'a'][0])
	}

	if (n%2 == 0 && maxCount > n/2) ||
		(n%2 == 1 && maxCount > (n/2+1)) {
		return ""
	}

	sort.Slice(count, func(i, j int) bool {
		return count[i][0] > count[j][0]
	})

	res := make([]byte, n)

	idx := 0

	for i := 0; count[i][0] > 0; i++ {
		b := byte('a' + count[i][1])

		for count[i][0] > 0 {
			res[idx] = b
			count[i][0]--
			idx = (idx + 2) % n
		}

	}

	return string(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
