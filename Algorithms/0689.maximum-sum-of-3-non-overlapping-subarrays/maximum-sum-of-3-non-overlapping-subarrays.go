package Problem0689

func maxSumOfThreeSubarrays(nums []int, k int) []int {
	n, maxsum := len(nums), 0
	sum := make([]int, n+1)
	posLeft := make([]int, n)
	posRight := make([]int, n)
	ans := make([]int, 3)
	for i := 0; i < n; i++ {
		sum[i+1] = sum[i] + nums[i]
	}
	// DP for starting index of the left max sum interval
	tot := sum[k] - sum[0]
	for i := k; i < n; i++ {
		if sum[i+1]-sum[i+1-k] > tot {
			posLeft[i] = i + 1 - k
			tot = sum[i+1] - sum[i+1-k]
		} else {
			posLeft[i] = posLeft[i-1]
		}
	}
	// DP for starting index of the right max sum interval
	// caution: the condition is ">= tot" for right interval, and "> tot" for left interval
	posRight[n-k] = n - k
	tot = sum[n] - sum[n-k]
	for i := n - k - 1; 0 <= i; i-- {
		if sum[i+k]-sum[i] >= tot {
			posRight[i] = i
			tot = sum[i+k] - sum[i]
		} else {
			posRight[i] = posRight[i+1]
		}
	}
	// test all possible middle interval
	for i := k; i <= n-2*k; i++ {
		l, r := posLeft[i-1], posRight[i+k]
		tot = (sum[i+k] - sum[i]) + (sum[l+k] - sum[l]) + (sum[r+k] - sum[r])
		if tot > maxsum {
			maxsum = tot
			ans[0] = l
			ans[1] = i
			ans[2] = r
		}
	}
	return ans
}
