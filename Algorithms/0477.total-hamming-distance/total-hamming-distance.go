package problem0477

func totalHammingDistance(nums []int) int {
	res := 0
	n := len(nums)

	for j := 0; j < 32; j++ {
		k := 0
		for i := 0; i < n; i++ {
			k += (nums[i] >> uint(j)) & 1
		}
		// 在第 j 个 bit 位上，
		// 共有 k 个数是 1 ， (n-k) 个位是 0
		// 那么，在第 j 个 bit 位上，会产生 k*(n-k) 个 hamming distance
		res += k * (n - k)
	}

	return res
}
