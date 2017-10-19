package Problem0321

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	size1 := len(nums1)
	size2 := len(nums2)

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
		return true
	}

	var dfs func(int, int, []int)
	dfs = func(i1, i2 int, temp []int) {
		if !isBigger(temp) {
			return
		}

		if len(temp) == k {
			copy(res, temp)
			return
		}

		for i := i1; i < size1 && size1+size2-i >= k-len(temp); i++ {
			dfs(i+1, i2, append(temp, nums1[i]))
		}

		for i := i2; i < size2 && size1+size2-i >= k-len(temp); i++ {
			dfs(i1, i+1, append(temp, nums2[i]))
		}
	}

	dfs(0, 0, temp)

	return res
}
