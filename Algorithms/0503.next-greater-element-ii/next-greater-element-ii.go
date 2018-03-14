package problem0503

func nextGreaterElements(nums []int) []int {
	size := len(nums)

	res := make([]int, size)
	for i := range res {
		res[i] = -1
	}

	// stack 中存放的是 nums 的递减子序列的索引号，例如
	// nums = {9,8,7,3,2,1,6}
	// 在遇到 6 以前，stack = {9,8,7,3,2,1} 的索引号
	// 在遇到 6 以后，stack = {9,8,7,6} 的索引号，
	// 此次已经得知了，{3,2,1} 的 nextGreater 是 6
	stack := make([]int, 0, size)

	// i < size*2 是为了避免 stack 中留有的最大值无法取出导致的无限循环。因为
	// 当 i < size 时，如果存在 nums[i] < nums[j%size]， 定有 j < i+size < size*2
	// size*2 内找不到，就永远也找不到了
	for i := 0; i < size*2; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i%size] {
			res[stack[len(stack)-1]] = nums[i%size]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%size)
	}

	return res
}
