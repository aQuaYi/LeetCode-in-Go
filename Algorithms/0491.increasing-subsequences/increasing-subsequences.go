package Problem0491

import (
	"fmt"
)

func findSubsequences(nums []int) [][]int {
	size := len(nums)
	temp := make([][]int, 1, (size-1)*(size-1))
	temp[0] = []int{}

	for _, n := range nums {
		sizeRes := len(temp)
		for i := 0; i < sizeRes; i++ {
			if len(temp[i]) == 0 ||
				temp[i][len(temp[i])-1] <= n {
				t := make([]int, len(temp[i])+1)
				copy(t, temp[i])
				t[len(temp[i])] = n
				temp = append(temp, t)
			}
		}
	}

	isAdded := make(map[string]bool)
	res := make([][]int, 0, len(temp))
	for _, t := range temp {
		key := fmt.Sprint(t)
		if len(t) >= 2 && !isAdded[key] {
			res = append(res, t)
			isAdded[key] = true
		}
	}

	return res
}
