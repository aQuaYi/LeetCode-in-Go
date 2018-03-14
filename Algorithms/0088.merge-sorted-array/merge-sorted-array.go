package problem0088

// 本题的要求是，把nums1的前m项和nums2的前n项合并，放入nums1中。
func merge(nums1 []int, m int, nums2 []int, n int) {
	// 深度复制 nums1
	temp := make([]int, m)
	copy(temp, nums1)

	j, k := 0, 0
	for i := 0; i < len(nums1); i++ {
		// nums2用完了
		if k >= n {
			nums1[i] = temp[j]
			j++
			continue
		}
		// temp 用完了
		if j >= m {
			nums1[i] = nums2[k]
			k++
			continue
		}
		// 比较后，放入
		if temp[j] < nums2[k] {
			nums1[i] = temp[j]
			j++
		} else {
			nums1[i] = nums2[k]
			k++
		}
	}
}
