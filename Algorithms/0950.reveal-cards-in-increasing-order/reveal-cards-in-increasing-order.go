package problem0950

import "sort"

func deckRevealedIncreasing(deck []int) []int {
	sort.Ints(deck)

	size := len(deck)

	t := make([]int, 0, size*2)
	b := 0

	for i := size - 1; i >= 0; i-- {
		t = append(t, deck[i])
		if i > 0 {
			t = append(t, t[b])
			b++
		}
	}

	for i, j := b, len(t)-1; i < j; i, j = i+1, j-1 {
		t[i], t[j] = t[j], t[i]
	}

	return t[b:]
}
