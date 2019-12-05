package problem0342

func isPowerOfFour(num int) bool {
	if num <= 0 {
		return false
	}

	// 如果 num 不是2的N次方, 那么也不是4的N次方
	if num&(num-1) != 0 {
		return false
	}

	// 过滤
	if num&0x55555555 == 0 {
		return false
	}

	return true
}
