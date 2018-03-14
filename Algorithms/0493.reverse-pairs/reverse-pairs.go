package problem0493

func reversePairs(nums []int) int {
	return count(nums, 0, len(nums))
}

func count(nums []int, beg, end int) int {
	if beg+1 >= end {
		// 当 nums[beg:end] 中只有 1 或 0 个 元素时
		// 直接返回 0
		return 0
	}

	// 把 nums[beg:end] 按照 mid 切分成两段
	// 分别统计两段中符合题意的元素对
	mid := beg + (end-beg)/2
	res := count(nums, beg, mid) + count(nums, mid, end)

	i, j := beg, mid
	for i < mid && j < end {
		if nums[i] > 2*nums[j] {
			// nums[i] 是 nums[i:mid] 中最小的数
			// nums[i] > 2*nums[j] 的话，
			// nums[i:mid] 中的任意一个数和 nums[j] 都是满足题意的 pair
			// 所以，一次就找到了 mid - i 对 pair
			res += mid - i
			j++
		} else {
			i++
		}
	}

	copy(nums[beg:end], merge(nums[beg:mid], nums[mid:end]))

	return res
}

// a, b 都是升序的切片
// merge 把 a，b 合并起来，并保持升序
func merge(a, b []int) []int {
	lenA, lenB := len(a), len(b)
	res := make([]int, lenA+lenB)

	var i, j, k int
	for i < lenA && j < lenB {
		if a[i] < b[j] {
			res[k] = a[i]
			i++
		} else {
			res[k] = b[j]
			j++
		}
		k++
	}

	if i == lenA {
		copy(res[k:], b[j:])
	} else {
		copy(res[k:], a[i:])
	}

	return res
}
