package Problem0137

// 假设 nums 中的数字为 32 bit
func singleNumber(nums []int) int {
	n := len(nums)
	count := make([]int, 32)
	res := 0

	for i := 0; i < 32; i++ {
		for j := 0; j < n; j++ {
			count[i] += ((nums[j] >> i) & 1) //首先把输入数字的第i位加起来。
		}
		count[i] = count[i] % 3 //然后求它们除以3的余数。

		res |= (count[i] << i) //把二进制表示的结果转化为十进制表示的结果
	}

	return res
}
