package problem0970

import (
	"sort"
)

func powerfulIntegers(x int, y int, bound int) []int {
	res := combine(powers(x, bound),
		powers(y, bound),
		bound)
	return removeRepeated(res)
}

func powers(n, bound int) []int {
	res := make([]int, 1, 128)
	res[0] = 1
	if n == 1 {
		return res
	}
	for i := n; i < bound; i *= n {
		res = append(res, i)
	}
	return res
}

func combine(ax, ay []int, bound int) []int {
	res := make([]int, 0, 128)
	for _, m := range ax {
		for _, n := range ay {
			s := m + n
			if s > bound {
				break
			}
			res = append(res, s)
		}
	}
	return res
}

func removeRepeated(nums []int) []int {
	sort.Ints(nums)

	size := len(nums)

	last, j := -1, -1
	for i := 0; i < size; i++ {
		if last == nums[i] { // nums[i]>0 for any i
			continue
		}
		j++
		nums[j], last = nums[i], nums[i]
	}
	return nums[:j+1]
}
