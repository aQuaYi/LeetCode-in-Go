package Problem0368

import "sort"

func largestDivisibleSubset(a []int) []int {
	var res []int

	size := len(a)

	sort.Ints(a)

	var i, j, k int

	for i = 0; i < size; i++ {
		temp := make([]int, 0, size-i)
		temp = append(temp, a[i])
		for j = i + 1; j < size; j++ {
			ok := true
			for k = 0; k < len(temp); k++ {
				if a[j]%temp[k] != 0 {
					ok = false
					break
				}
			}
			if ok {
				temp = append(temp, a[j])
			}
		}
		if len(res) < len(temp) {
			res = temp
		}
	}

	return res
}
