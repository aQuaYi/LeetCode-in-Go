package problem0950

import "sort"

func deckRevealedIncreasing(deck []int) []int {
	sort.Ints(deck)

	size := len(deck)
	t := make([]int, 1, size*2)

	t[0] = deck[size-1]

	b := 0
	for i := size - 2; i > 0; i-- {
		t = append(t, deck[i])
		t = append(t, t[b])
		b++
	}

	t = append(t, deck[0])

	for i, j := b, len(t)-1; i < j; i, j = i+1, j-1 {
		t[i], t[j] = t[j], t[i]
	}

	return t[b:]
}
