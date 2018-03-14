package problem0047

import "sort"

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)

	n := len(nums)
	// vector 是一组可能的解答
	vector := make([]int, n)
	// taken[i] == true 表示 nums[i] 已经出现在了 vector 中
	taken := make([]bool, n)

	var ans [][]int

	makePermutation(0, n, nums, vector, taken, &ans)

	return ans
}

// 这个方式和我原来的方式相比，
// 增加了比较的次数
// 但是，减少了切片生成的次数。
// 所以，速度会快一些。
func makePermutation(cur, n int, nums, vector []int, taken []bool, ans *[][]int) {
	if cur == n {
		tmp := make([]int, n)
		copy(tmp, vector)
		*ans = append(*ans, tmp)
		return
	}

	used := make(map[int]bool, n-cur)

	for i := 0; i < n; i++ {

		if !taken[i] && !used[nums[i]] {
			used[nums[i]] = true

			// 准备使用 nums[i]，所以，taken[i] == true
			taken[i] = true
			// NOTICE: 是 vector[cur]
			vector[cur] = nums[i]

			makePermutation(cur+1, n, nums, vector, taken, ans)

			// 下一个循环中
			// vector[cur] = nums[i+1]
			// 所以，在这个循环中，恢复 nums[i] 自由
			taken[i] = false
		}
	}
}
