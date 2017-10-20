package Problem0316

import (
	"sort"
	"strings"
)

func removeDuplicateLetters(s string) string {
	var i, j, k, l int
	// rec 记录字母每次出现的位置
	rec := make(map[byte][]int, 26)
	for i = range s {
		rec[s[i]] = append(rec[s[i]], i)
	}
	base := byte(255)
	i, j, k = -1, 0, 0
	for ; k < len(s); k++ {
		if len(s[k]) == 1 {
			j = k
			base = s[k]

		}
	}
	// bytes 出现过的字母
	bytes := make([]int, 0, len(rec))
	for k := range rec {
		bytes = append(bytes, k)
	}
	// 字母升序
	sort.Ints(bytes)

	var k, bi, bj, base int
	for i, bi = range bytes {
		base = rec[bi][0]
		for _, bj = range bytes[i+1:] {
			k = 0
			for k+1 < len(rec[bj]) && rec[bj][k] < base {
				k++
			}
			if k > 0 {
				rec[bj] = rec[bj][k:]
			}
		}
	}

	res := make([]string, len(s))
	for b, idxs := range rec {
		res[idxs[0]] = string(b + 'a')
	}

	return strings.Join(res, "")
}
