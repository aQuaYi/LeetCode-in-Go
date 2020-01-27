package problem0004

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func findKth(k int, nums1 []int, i int, nums2 []int, j int) int {
	if i > j {
		return findKth(k, nums2, j, nums1, i)
	} else if i == 0 {
		return nums2[k-1]
	} else if k == 1 {
		return min(nums1[0], nums2[0])
	} else {
		var pa = min(k/2, i)
		var pb = k - pa
		if nums1[pa-1] < nums2[pb-1] {
			return findKth(k-pa, nums1[pa:], i-pa, nums2, j)
		} else if nums1[pa-1] > nums2[pb-1] {
			return findKth(k-pb, nums1, i, nums2[pb:], j-pb)
		} else {
			return nums1[pa-1]
		}
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var l1, l2 = len(nums1), len(nums2)
	if (l1+l2)%2 == 1 {
		return float64(findKth((l1+l2)/2+1, nums1, l1, nums2, l2))
	}
	return float64(findKth((l1+l2)/2, nums1, l1, nums2, l2)+findKth((l1+l2)/2+1, nums1, l1, nums2, l2)) / 2
}
