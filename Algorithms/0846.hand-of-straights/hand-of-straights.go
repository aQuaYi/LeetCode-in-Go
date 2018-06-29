package problem0846

import (
	"sort"
)

func isNStraightHand(hand []int, W int) bool {
	if W == 1 {
		return true
	}

	if W > len(hand) || len(hand)%W != 0 {
		return false
	}

	size := len(hand) / W

	groups := make([][]int, size)

	sort.Ints(hand)

	for _, c := range hand {
		i := 0
		for ; i < size; i++ {
			if len(groups[i]) == W {
				continue
			}
			last := len(groups[i]) - 1
			if last == -1 || groups[i][last]+1 == c {
				groups[i] = append(groups[i], c)
				break
			}
		}
		if i == size {
			return false
		}
	}

	return true
}
