package problem0875

import (
	"math"
)

func minEatingSpeed(piles []int, h int) int {
	sum, r := sumAndMax(piles)
	l := (sum + h - 1) / h

	for l < r {
		m := (l + r) / 2
		if canEatAll(m, h, piles) {
			r = m
		} else {
			l = m + 1
		}
	}

	return r
}

func sumAndMax(a []int) (int, int) {
	sum := 0
	mx := 0
	for _, n := range a {
		sum += n
		mx = max(mx, n)
	}
	return sum, mx
}

func canEatAll(k, h int, piles []int) bool {
	r := 1 / float64(k)
	for _, p := range piles {
		fp := float64(p)
		h -= int(math.Ceil(fp * r))
		if h < 0 {
			return false
		}
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
