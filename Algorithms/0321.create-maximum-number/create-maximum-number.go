package problem0321

func maxNumber(nums1, nums2 []int, k int) []int {
	m, n := len(nums1), len(nums2)

	res := make([]int, k)
	var temp []int

	for i := max(0, k-n); i <= m && i <= k; i++ {
		temp = combine(choose(i, nums1), choose(k-i, nums2))
		if isBigger(temp, res) {
			copy(res, temp)
		}
	}

	return res
}

// 从 nums 当中挑选 k 个数，其组成的 []int 最大
// 不满足 0 <= k <= len(nums) 会 panic
func choose(k int, nums []int) []int {
	if k == len(nums) {
		return nums
	}

	res := make([]int, k)
	// idx 是 res 上次获取的 nums 的值的索引号
	idx := -1
	for i := 0; i < k; i++ {
		idx++
		// res[i] 是 nums[idx:len(nums)-k+i+1] 中的最大值
		res[i] = nums[idx]
		for j := idx + 1; j <= len(nums)-k+i; j++ {
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
