package Problem0016

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {

	// 排序后，可以按规律查找
	sort.Ints(nums)
	res, delta := 0, math.MaxInt64

	for i := range nums {
		// 避免重复计算
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1

		for l < r {
			s := nums[i] + nums[l] + nums[r]
			temp := abs(target - s)
			switch {
			case temp < delta:

				l++
			case s > 0:
				// 较大的 r 需要变小
				r--
			default:
				res = append(res, []int{nums[i], nums[l], nums[r]})
				// 为避免重复添加，l 和 r 都需要移动到不同的元素上。
				l, r = next(nums, l, r)
			}
		}
	}

	return res
}

func next(nums []int, l, r int) (int, int) {
	for l < r {
		switch {
		case nums[l] == nums[l+1]:
			l++
		case nums[r] == nums[r-1]:
			r--
		default:
			l++
			r--
			return l, r
		}
	}

	return l, r
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
