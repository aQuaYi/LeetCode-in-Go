package problem0004

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := combine(nums1, nums2)
	return medianOf(nums)
}

func combine(mis, njs []int) []int {
	lenMis, i := len(mis), 0
	lenNjs, j := len(njs), 0
	res := make([]int, lenMis+lenNjs)

	for k := 0; k < lenMis+lenNjs; k++ {
		if i == lenMis ||
			(i < lenMis && j < lenNjs && mis[i] > njs[j]) {
			res[k] = njs[j]
			j++
			continue
		}

		if j == lenNjs ||
			(i < lenMis && j < lenNjs && mis[i] <= njs[j]) {
			res[k] = mis[i]
			i++
		}
	}

	return res
}

func medianOf(nums []int) float64 {
	l := len(nums)

	if l == 0 {
		panic("切片的长度为0，无法求解中位数。")
	}

	if l%2 == 0 {
		return float64(nums[l/2]+nums[l/2-1]) / 2.0
	}

	return float64(nums[l/2])
}
