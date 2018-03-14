package problem0672

func flipLights(n, m int) int {
	if m == 0 {
		return 1
	}

	switch n {
	// 当只有一盏灯的时候
	// 如果只按 1 次：
	// 	按键1,3,4，会让灯熄灭
	// 	按键2，会让灯保持点亮
	// 	所以，一共是 2 种状态
	// 如果可以按超过 1 次，那么灯的状态，只会比按 1 次多，
	// 	但是，一盏灯最多只会有 2 种状态
	// 所以，n == 1， m>0 时，返回 2
	case 1:
		return 2
		// 与 n == 1 的情况类似
	case 2:
		if m == 1 {
			return 3
		}
		return 4
	// 当 n >= 3 时
	// 根据 4 个按钮控制灯泡的位置不同，可以发现
	// 所有 3*k+1 (k=0,1,2,...) 位置上的灯泡状态总是相同的
	// 3*k+2, 3*k+3 (k=0,1,2,...) 也是如此
	// 所以，
	// 可以把所有的灯泡，每 3 个分成一组
	// n >=3 被简化成了 n == 3
	//
	// 在纸上演算一下，可能出现的各种情况即可
	default:
		switch m {
		case 1:
			return 4
		case 2:
			return 7
		default:
			return 8
		}
	}
}
