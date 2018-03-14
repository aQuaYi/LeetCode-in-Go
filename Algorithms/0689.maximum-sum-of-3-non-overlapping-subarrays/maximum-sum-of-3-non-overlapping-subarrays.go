package problem0689

func maxSumOfThreeSubarrays(nums []int, k int) []int {
	n := len(nums) - k + 1

	// sum[i] = sum(nums[i:i+k])
	sum := make([]int, n)
	sumK := 0
	for i := 0; i < len(nums); i++ {
		sumK += nums[i]
		if i >= k {
			sumK -= nums[i-k]
		}
		if i >= k-1 {
			sum[i-k+1] = sumK
		}
	}

	// left[i] == j 表示，在 sum[:i+1] 中，最大值的索引号为 j
	left := make([]int, n)
	indexOfMax := 0
	for i := 0; i < n; i++ {
		if sum[indexOfMax] < sum[i] {
			indexOfMax = i
		}
		left[i] = indexOfMax
	}

	indexOfMax = n - 1
	// right[i] == j 表示，在 sum[i:] 中，最大值的索引号为 j
	right := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		if sum[indexOfMax] < sum[i] {
			indexOfMax = i
		}
		right[i] = indexOfMax
	}

	a := []int{0, k, k + k}
	for y := k ; y < n-k; y++ {
		x, z := left[y-k], right[y+k]
		if sum[a[0]]+sum[a[1]]+sum[a[2]] < sum[x]+sum[y]+sum[z] {
			a[0], a[1], a[2] = x, y, z
		}
	}

	return a
}
