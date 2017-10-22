package Problem0368

import "sort"

func largestDivisibleSubset(a []int) []int {
	size := len(a)
	sort.Ints(a)

	temps := make([][]int, 1, size-1)
	temps[0] = []int{a[0]}

	var i, j, k, t int
	for i = 1; i < size; i++ {
		for _, temp := range temps {
			for j, t = range temp {
				if a[i]%t != 0 {
					break
				}
			}

		}
	}

	var res []int

	return res
}
