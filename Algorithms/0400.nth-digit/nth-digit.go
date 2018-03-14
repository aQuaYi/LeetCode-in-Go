package problem0400

func findNthDigit(n int) int {
	// step 1: 寻找拥有 NthDigit 的数的位数 digits
	// 有 count 个数 是 digits 位数
	count, digits := 9, 1
	// num 是最小的 digits 位数
	num := 1
	for n-count*digits > 0 {
		n -= count * digits
		count *= 10
		digits++
		num *= 10
	}

	// step 2: 找到拥有 NthDigit 的数
	// index 是 NthDigit 是目标数中的 索引号
	index := n % digits
	if index == 0 {
		index = digits
	}
	// 让 num 成为拥有 NthDigit 的数
	num += n / digits
	if index == digits {
		num--
	}

	// step 3: 找到 NthDigit
	for i := index; i < digits; i++ {
		num /= 10
	}
	return num % 10
}
