package Problem0321

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	size1 := len(nums1)
	size2 := len(nums2)

	if size1 > size2 {
		return maxNumber(nums2, nums1, k)
	}

	res := make([]int, k)
	temp := make([]int, 0, k)

	isBigger := func(temp []int) bool {
		for i := range temp {
			if temp[i] > res[i] {
				return true
			} else if temp[i] < res[i] {
				return false
			}
		}
		return false
	}

	for i := 0; i < size1 && size1+size2-i >= k; i++ {
		temp = combine(outOf(nums1, size1-i), outOf(nums2, k-(size1-i)))
		if isBigger(temp) {
			copy(res, temp)
		}
	}

	return res
}

// 从 nums 当中挑选 k 个数，其组成的 []int 最大
func outOf(nums []int, k int) []int {
	res := make([]int, k)
	var idx = -1
	for i := 0; i < k; i++ {
		idx++
		res[i] = nums[idx]
		for j := idx; j <= len(nums)-k+i; j++ {
			if res[i] < nums[j] {
				res[i] = nums[j]
				idx = j
			}
		}
	}
	return res
}

// 混合 nums1 和 nums2 使得其组成的 []int 最大
func combine(nums1, nums2 []int) []int {
	size1 := len(nums1)
	size2 := len(nums2)
	res := make([]int, 0, size1+size2)

	var i, j int
	for i < size1 && j < size2 {
		if nums1[i] > nums2[j] {
			res = append(res, nums1[i])
			i++
		} else {
			res = append(res, nums2[j])
			j++
		}
	}

	res = append(res, nums1[i:]...)
	res = append(res, nums2[j:]...)

	return res
}
