package problem0137

//ref: https://cloud.tencent.com/developer/article/1131945

// 如果是出现两次的话，用一个 bit 就可以
// 比如 ones，初始为0
// 当 5 第 1 次出现， ones = 5
// 当 5 第 2 次出现， ones 清空为 0，表示 ones 可以去处理其他数字了
// 所以，最后 如果 ones != 0的话， ones 记录的就是只出现了一次的那个数字

// 公式是 ones = ones xor i

// 那么，如果是三次的话，香农定理，需要用 2 bits 进行记录

// 当 5 第 1 次出现的时候，ones = 5, twos = 0,  ones 记录这个数字
// 当 5 第 2 次出现的时候，ones = 0, twos = 5， twos 记录了这个数字
// 当 5 第 3 次出现的时候，ones = 0, twos = 0， 都清空了，可以去处理其他数字了
// 所以，如果有某个数字出现了 1 次，就存在 ones 中，出现了 2 次，就存在 twos 中

// 公式方面， 上面 2 次的时候，ones 清空的公式是 ones = ones xor i
// 而第 3 次时， ones 要等于零， 而这时 twos 是 True ， 所以再 & 一个 twos 的非就可以， ones = (ones xor i) & ~twos
// 所以，总的公式是
//    ones = (ones xor i) & ~twos
//    twos = (twos xor i) & ~ones

func singleNumber(nums []int) int {
	ones, twos := 0, 0
	for i := 0; i < len(nums); i++ {
		ones = (ones ^ nums[i]) & ^twos
		twos = (twos ^ nums[i]) & ^ones
	}
	return ones
}
