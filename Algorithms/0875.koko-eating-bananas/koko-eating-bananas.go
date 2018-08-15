package problem0875

import "math"

func minEatingSpeed(piles []int, H int) int {
	sum := sumOf(piles)

	l := (sum + H - 1) / H
	r := l * 2

	for l < r {
		m := (l + r) / 2
		if canEatAll(m, H, piles) {
			r = m
		} else {
			l = m + 1
		}
	}

	return r
}

func sumOf(a []int) int {
	res := 0
	for _, n := range a {
		res += n
	}
	return res
}

func canEatAll(k, h int, piles []int) bool {
	r := 1 / float64(k)
	for _, p := range piles {
		h -= int(math.Ceil(float64(p) * r))
		if h < 0 {
			return false
		}
	}
	return true
}
