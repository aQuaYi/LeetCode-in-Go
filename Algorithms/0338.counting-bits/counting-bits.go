package problem0338

func countBits(num int) []int {
	res := make([]int, num+1)
	for i := 1; i <= num; i++ {
		// i>>1 == i/2
		// i&1  == i%2
		// 只是 位运算 更快
		//
		// 观察以下三个数的二进制表示
		// 5 : 101
		// 10: 1010, 10>>1 == 5
		// 11: 1011, 11>>1 == 5
		// 10 的二进制表示，含有 1 的个数，可以由 5 的答案 + 10%2 计算
		// 11 同理
		res[i] = res[i>>1] + i&1
	}
	return res
}
