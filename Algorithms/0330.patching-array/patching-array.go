package Problem0330

import "sort"
import "fmt"

func minPatches(nums []int, n int) int {
	count := n
	isGetted := make([]bool, n+1)

	// 利用 nums 已有的元素，填充isGetted
	var dfs func(int, int)
	dfs = func(sum, idx int) {
		if idx == len(nums) {
			if 0 < sum && sum <= n && !isGetted[sum] {
				isGetted[sum] = true
				count--
			}
			return
		}
		dfs(sum, idx+1)
		dfs(sum+nums[idx], idx+1)
	}

	check := func(x int) {
		var i, j int
		for i = 0; i < len(nums) && nums[i] <= x; i++ {
			sum := 0
			for j = i; j < len(nums) && nums[j] <= x; j++ {
				sum += nums[j]
				if sum == x {
					return
				}
				if sum > x {
					break
				}
			}
		}
		nums = append(nums, x)
		count++
		sort.Ints(nums)
	}

	sort.Ints(nums)

	var i int
	for i = 1; i <= n; i++ {
		check(i)
	}

	fmt.Println(nums)
	return count
}
