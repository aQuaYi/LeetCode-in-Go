package problem0875

import (
	"sort"
)

func minEatingSpeed(piles []int, H int) int {
	sort.Ints(piles)

	n := len(piles)

	l, r := 1, piles[n-1]
	for l < r {
		m := (l + r) / 2
		if canEatAll(m, H, n, piles) {
			r = m
		} else {
			l = m + 1
		}
	}

	return r
}

func canEatAll(k, h, n int, piles []int) bool {
	t := sort.Search(n, func(i int) bool {
		return piles[i] > k
	})

	for b, i := t, n-1; b <= i && t <= h; i-- {
		t += (piles[i] + k - 1) / k
	}

	return t <= h
}
