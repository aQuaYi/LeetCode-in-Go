package problem0846

import (
	"sort"
)

func isNStraightHand(hand []int, W int) bool {
	size := len(hand)
	group := size / W
	if group*W != size {
		// 没有办法保证每组的个数是 W
		return false
	}

	groups := make([][]int, group)
	for i := range groups {
		groups[i] = make([]int, 1, W)
		groups[i][0] = -1
	}

	sort.Ints(hand)

	for _, c := range hand {
		g := 0
		for ; g < group; g++ {
			if len(groups[g]) == W && W != 1 {
				continue
			}
			if groups[g][0] == -1 {
				groups[g][0] = c
				break
			}
			last := len(groups[g]) - 1
			if groups[g][last]+1 == c {
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
