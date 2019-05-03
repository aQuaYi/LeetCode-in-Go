package problem0970

import (
	"math"
	"sort"
)

func powerfulIntegers(x int, y int, bound int) []int {
	res := make([]int, 0, 1024)

	logBound := math.Log(float64(bound))
	logX := math.Log(float64(x))
	logY := math.Log(float64(y))

	maxX := int(math.Pow(float64(x), math.Round(logBound/logX)))
	maxY := int(math.Pow(float64(y), math.Round(logBound/logY)))

	for i := maxX; i >= 1; i /= x {
		for j := maxY; j >= 1; j /= y {
			if i+j > bound {
				continue
			}
			res = append(res, i+j)
		}
	}

	res = removeRepeated(res)

	return res
}

func removeRepeated(nums []int) []int {
	sort.Ints(nums)
	res := make([]int, 1, len(nums))
	res[0] = nums[0]
	last := nums[0]
	for _, n := range nums {
		if n == last {
			continue
		}
		res = append(res, n)
		last = n
	}
	return res
}
