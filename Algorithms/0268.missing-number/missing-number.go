package problem0268

func missingNumber(nums []int) int {
	xor := 0

	for i, n := range nums {
		xor ^= i ^ n
	}

	// 所有的 i 再加上 len(nums)，就相当于 0,1,...,n。
	// 利用 相同的数异或为0，及其交换律
	// xor 最后的值，就是那个缺失的数

	return xor ^ len(nums)
}
