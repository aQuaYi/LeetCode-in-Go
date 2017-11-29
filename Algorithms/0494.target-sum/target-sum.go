package Problem0494

func findTargetSumWays(nums []int, S int) int {

	mid := len(nums) / 2
	left := make(map[int]int, len(nums))
	right := make(map[int]int, len(nums))
	dfs(nums[:mid], 0, 0, left)
	dfs(nums[mid:], 0, 0, right)

	res := 0
	for l, c := range left {
		res += c * right[S-l]
	}

	return res
}

func dfs(nums []int, index, sum int, rec map[int]int) {
	if index == len(nums) {
		rec[sum]++
		return
	}

	dfs(nums, index+1, sum+nums[index], rec)
	dfs(nums, index+1, sum-nums[index], rec)
}
