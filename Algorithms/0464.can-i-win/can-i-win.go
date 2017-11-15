package Problem0464

func helper(maxInt int, remains int, bit int, dp map[int]bool, bitmap []int) bool {
	if val, ok := dp[bit]; ok {
		return val
	}

	if remains <= maxInt {
		for i := maxInt; i >= remains; i-- {
			if (bit & bitmap[i]) == 0 {
				dp[bit] = true
				return true
			}
		}
	}

	for i := maxInt; i > 0; i-- {
		if (bit & bitmap[i]) > 0 {
			continue
		}
		bit |= bitmap[i]
		ret := helper(maxInt, remains-i, bit, dp, bitmap)
		bit &= (^bitmap[i])
		if ret == false {
			dp[bit] = true
			return true
		}
	}

	dp[bit] = false

	return false
}

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	dp := make(map[int]bool)
	sum := 0
	bitmap := make([]int, maxChoosableInteger+1)
	for i := maxChoosableInteger; i > 0; i-- {
		sum += i
		bitmap[i] = 1 << uint8(i)
	}

	// 这种情况下，双方都不能赢
	if sum < desiredTotal {
		return false
	}

	return helper(maxChoosableInteger, desiredTotal, 0, dp, bitmap)
}
