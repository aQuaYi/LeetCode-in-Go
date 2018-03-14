package problem0227

func calculate(s string) int {
	// n1 opt1 n2 opt2 n3
	// ↓   ↓   ↓   ↓   ↓
	// 1   +   2   *   3
	var n1, n2, n3 int
	var opt1, opt2 byte
	opt1 = '+'

	idx := 0
	nextN := func() int {
		n := 0
		// 跳过空格
		for idx < len(s) && s[idx] == ' ' {
			idx++
		}
		// 获取数值
		for idx < len(s) && '0' <= s[idx] && s[idx] <= '9' {
			n = n*10 + int(s[idx]-'0')
			idx++
		}
		return n
	}

	nextOpt := func() byte {
		// opt 默认为 +，与 nextN() 的默认值为 0 配合
		// 当 s 的结尾为 空格 的时候，程序会进行一次 +0 的运算，但这不会影响结果。
		opt := byte('+')
		// 跳过空格
		for idx < len(s) && s[idx] == ' ' {
			idx++
		}
		// 获取操作符
		if idx < len(s) {
			opt = s[idx]
			idx++
		}
		return opt
	}

	n2 = nextN()
	for idx < len(s) {
		opt2 = nextOpt()
		n3 = nextN()

		if opt2 == '*' || opt2 == '/' {
			// 先乘除
			n2 = operate(n2, n3, opt2)
		} else {
			// 后加减
			n1 = operate(n1, n2, opt1)
			opt1 = opt2
			n2 = n3
		}
	}

	return operate(n1, n2, opt1)
}

func operate(a, b int, opt byte) int {
	switch opt {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	default: // '/'
		return a / b
	}
}
