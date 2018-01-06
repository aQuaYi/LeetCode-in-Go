package Problem0525

func findMaxLength(nums []int) int {
	size := len(nums)
	sums := make([]int, size+1)
	for i := 1; i < len(sums); i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}

	max := 0
	for j := len(sums) - 1; 0 <= j; j-- {
		for i := 0; i+max < j; i++ {
			if j-i == (sums[j]-sums[i])*2 {
				max = j - i
			}
		}
	}

	return max
}
