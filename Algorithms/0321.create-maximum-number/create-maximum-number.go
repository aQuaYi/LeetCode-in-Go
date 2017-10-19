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
		return false
	}

	for i := 0; i < size1 && i <= k; i++ {

	}

	return res
}
