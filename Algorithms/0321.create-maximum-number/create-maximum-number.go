package Problem0321

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	size1 := len(nums1)
	size2 := len(nums2)

	res := make([]int, k)
	var temp, temp1, temp2 []int

	var i int
	for i = 0; i <= size1 && i <= k; i++ {
		if size2 < k-i { // nums2 不够长
			continue
		}
		temp1 = outOf(nums1, i)
		temp2 = outOf(nums2, k-i)
		temp = combine(temp1, temp2)
		if isBigger(temp, res) {
			copy(res, temp)
		}
	}

	return res
}

// 从 nums 当中挑选 k 个数，其组成的 []int 最大
func outOf(nums []int, k int) []int {
	if k == 0 {
		return []int{}
	}

	if k == len(nums) {
		return nums
	}

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
	if isBigger(nums1, nums2) {
		return combine(nums2, nums1)
	}

	size1 := len(nums1)
	size2 := len(nums2)
	res := make([]int, 0, size1+size2)

	var i, j int
	for i < size1 && j < size2 {
		if nums1[i] > nums2[j] || isBigger(nums1[i:], nums2[j:]) {
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

func isBigger(a1, a2 []int) bool {
	s1 := len(a1)
	s2 := len(a2)
	if s1 > s2 {
		return !isBigger(a2, a1)
	}

	for i := 0; i < s1; i++ {
		if a1[i] > a2[i] {
			return true
		} else if a1[i] < a2[i] {
			return false
		}
	}

	return false
}
