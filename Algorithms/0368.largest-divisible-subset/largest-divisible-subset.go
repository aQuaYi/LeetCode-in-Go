package Problem0368

import "sort"

func largestDivisibleSubset(a []int) []int {
	size := len(a)
	sort.Ints(a)
	res := make([]int, 0, size)

	if a[0] == 1 {
		res = append(res, 1)
		a = a[1:]
	}

	temps := make([][]int, 0, size)
	for len(a) > 0 {
		t := a[0]
		temp := make([]int, 1, len(a))
		tempA := make([]int, 0, len(a))
		temp[0] = t
		for i := 1; i < len(a); i++ {
			if a[i]%t == 0 {
				temp = append(temp, a[i])
			} else {
				tempA = append(tempA, a[i])
			}
		}
	}

	return res
}
