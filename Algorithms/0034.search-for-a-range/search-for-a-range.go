package problem0034

func searchRange(nums []int, target int) []int {
	// 查看target是否存在与nums中
	index := search(nums, target)
	if index == -1 {
		return []int{-1, -1}
	}

	// 利用二分法，查找第一个target
	first := index
	for {
		f := search(nums[:first], target)
		if f == -1 {
			break
		}
		first = f
	}

	// 利用二分法，查找最后一个target
	last := index
	for {
		l := search(nums[last+1:], target)
		if l == -1 {
			break
		}
		// 注意这里与查找first的不同
		last += l + 1
	}

	return []int{first, last}
}

// 二分查找法
func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	var median int

	for low <= high {
		median = (low + high) / 2

		switch {
		case nums[median] < target:
			low = median + 1
		case nums[median] > target:
			high = median - 1
		default:
			return median
		}
	}
	return -1
}
