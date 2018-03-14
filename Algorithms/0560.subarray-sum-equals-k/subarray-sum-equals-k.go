package problem0560

func subarraySum(a []int, k int) int {
	res, sum := 0, 0
	rec := make(map[int]int, len(a))
	rec[0] = 1

	for i := range a {
		// sum 等于 a[:i] 中所有元素之和
		sum += a[i]
		// rec[x] 是 a[:i] 中所有和为 x 的连续子序列的个数
		// 假设 rec[sum-k] == 1
		// 那么，有且仅有一个 j，满足 0<= j < i，且 a[j:i] 中所有元素的和为 k
		res += rec[sum-k]
		// 更新 rec
		rec[sum]++
	}

	return res
}
