package problem0846

import (
	"sort"
)

func isNStraightHand(hand []int, W int) bool {
	if W == 1 {
		return true
	}

	size := len(hand)
	group := size / W
	if group*W != size {
		// 没有办法保证每组的个数是 W
		return false
	}

	groups := make([][]int, group)

	sort.Ints(hand)

	for _, c := range hand {
		g := 0
		for ; g < group; g++ {
			if len(groups[g]) == W {
				continue
			}
			last := len(groups[g]) - 1
			if last == -1 || groups[g][last]+1 == c {
				groups[g] = append(groups[g], c)
				break
			}
		}

		if g == group {
			return false
		}

	}

	return true
}
