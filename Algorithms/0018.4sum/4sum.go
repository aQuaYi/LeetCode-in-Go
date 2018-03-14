package problem0018

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(nums)

	for i := 0; i < len(nums)-3; i++ {
		// 避免添加重复的结果
		// i>0 是为了防止nums[i-1]溢出
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums)-2; j++ {
			// 避免添加重复的结果
			// nums[j-1]虽然不会溢出
			// 但是不添加 j > i+1 的话，会漏掉那些i和j为相同数值的答案
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			l, r := j+1, len(nums)-1
			for l < r {
				s := nums[i] + nums[j] + nums[l] + nums[r]
				switch {
				case s < target:
					l++
				case s > target:
					r--
				default:
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					l, r = next(nums, l, r)
				}
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
