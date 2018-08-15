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
	for i := 0; i < len(piles) && h >= 0; i++ {
		p := float64(piles[i])
		h -= int(math.Ceil(p * r))
	}
	return h >= 0
}
