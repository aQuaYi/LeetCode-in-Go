package Problem0493

func reversePairs(nums []int) int {
	return count(nums, 0, len(nums))
}

func count(nums []int, beg, end int) int {
	if beg+1 == end {
		return 0
	}

	mid := beg + (end-beg)/2

	res := count(nums, beg, mid) + count(nums, mid, end)

	i, j := beg, mid

	for i < mid && j < end {
		if nums[i] > 2*nums[j] {
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
