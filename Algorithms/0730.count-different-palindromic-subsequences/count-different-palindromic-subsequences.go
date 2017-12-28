package Problem0730

import (
	"strings"
)

const mod = 1000000007

var elements = [4]string{"a", "b", "c", "d"}

func countPalindromicSubsequences(s string) int {
	check := make(map[[2]int]int, 1024)
	var dfs func(int, int) int

	cache := func(start, end int) int {
		if end <= start+2 {
			return end - start
		}
		key := [2]int{start, end}
		res, ok := check[key]
		if ok {
			return res
		}

		check[key] = dfs(start, end)

		return check[key]
	}

	dfs = func(start, end int) int {
		count := 0
		segment := s[start:end]

		for _, x := range elements {
			i := strings.Index(segment, x)
			j := strings.LastIndex(segment, x)
			if i == -1 {
				continue
			}
			if i == j {
				count++
			} else {
				count += cache(start+i+1, start+j) + 2
			}
		}

		return count % mod
	}

	return cache(0, len(s))
}
