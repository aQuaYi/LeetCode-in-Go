package problem0401

import (
	"fmt"
	"sort"
)

func readBinaryWatch(n int) []string {
	res := make([]string, 0, 8)
	leds := make([]bool, 10)

	var dfs func(int, int)
	dfs = func(idx, n int) {
		var h, m int
		if n == 0 {
			m, h = get(leds[:6]), get(leds[6:])
			if h < 12 && m < 60 {
				res = append(res, fmt.Sprintf("%d:%02d", h, m))
			}
			return
		}

		for i := idx; i < len(leds)-n+1; i++ {
			leds[i] = true
			dfs(i+1, n-1)
			leds[i] = false
		}
	}

	dfs(0, n)

	sort.Strings(res)

	return res
}

var bs = []int{1, 2, 4, 8, 16, 32}

func get(leds []bool) int {
	var sum int
	size := len(leds)
	for i := 0; i < size; i++ {
		if leds[i] {
			sum += bs[i]
		}
	}

	return sum
}
