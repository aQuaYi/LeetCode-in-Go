package Problem0400

func findNthDigit(n int) int {
	// 寻找拥有 NthDigit 的数的位数 digits
	// 有 9 个数 (base) 是 1 位数 (digits)
	base, digits := 9, 1
	num := 1
	for n-base*digits > 0 {
		n -= base * digits
		base *= 10
		digits++
		num *= 10
	}

	// index 是 NthDigit 是目标数中的 索引号
	index := n % digits
	if index == 0 {
		index = digits
	}

	// // 先让 num 成为最小的 digits 位数
	// num := 1
	// for i := 1; i < digits; i++ {
	// 	num *= 10
	// }

	// 再让 num 成为拥有 NthDigit 的数
	num += n / digits
	if index == digits {
		num--
	}

	// 找到 NthDigit
	for i := index; i < digits; i++ {
		num /= 10
	}
	return num % 10
}
