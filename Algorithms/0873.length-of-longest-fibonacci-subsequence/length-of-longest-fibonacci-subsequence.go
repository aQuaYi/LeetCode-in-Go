package problem0873

import "sort"

func lenLongestFibSubseq(a []int) int {
	size := len(a)
	res := 0

	for k := 2; k < size; k++ {
		i, j := 0, k-1
		for i < j {
			s := a[i] + a[j]

			if s < a[k] {
				i++
				continue
			}

			if a[k] < s {
				j--
				continue
			}

			count := 3
			x0, x1 := a[j], a[k]
			idx := k
			for {
				x2 := x0 + x1
				idx += sort.SearchInts(a[idx:], x2)
				if idx == size || a[idx] != x2 {
					break
				}
				count++
				x0, x1 = x1, x2
			}

			res = max(res, count)
			i++
			j--
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
