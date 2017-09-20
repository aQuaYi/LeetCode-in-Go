package Problem0137

// 假设 nums 中的数字为 32 bit
func singleNumber(nums []int) int {
	n := len(nums)
	res := 0
	base := 1

	var i uint
	for ; i < 32; i++ {
		temp := 0
		for j := 0; j < n; j++ {
			temp += nums[j] % 2 //首先把输入数字的第i位加起来。
			nums[j] /= 2
		}
		temp %= 3          //然后求它们除以3的余数。
		res += temp * base //把二进制表示的结果转化为十进制表示的结果
		base *= 2
	}

	return res
}
