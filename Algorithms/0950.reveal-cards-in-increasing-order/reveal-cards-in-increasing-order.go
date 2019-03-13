package problem0950

import "sort"

func deckRevealedIncreasing(deck []int) []int {
	sort.Ints(deck)

	size := len(deck)
	res := make([]int, size)

	b, i, jump := 0, 0, 2
	for _, c := range deck {
		if i >= size {
			b += jump / 2
			jump *= 2
			i = b
		}
		res[i] = c
		i += jump
	}

	return res
}
