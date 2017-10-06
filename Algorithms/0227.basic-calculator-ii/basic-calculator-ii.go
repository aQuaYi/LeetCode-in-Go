package Problem0227

func calculate(s string) int {
	var n1, n2, n3 int
	var opt1, opt2 byte
	opt1 = '+'

	idx := 0
	nextN := func() int {
		n := 0

		for idx < len(s) && s[idx] == ' ' {
			idx++
		}

		for idx < len(s) && '0' <= s[idx] && s[idx] <= '9' {
			n = n*10 + int(s[idx]-'0')
			idx++
		}

		return n
	}

	nextOpt := func() byte {
		opt := byte('+')

		for idx < len(s) && s[idx] == ' ' {
			idx++
		}

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
			n2 = operate(n2, n3, opt2)
		} else {
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
	default:
		return a / b
	}
}
