package Problem0224

// s 中不会含有 × ÷，所以，可以不用理会括号
func calculate(s string) int {
	num := 0
	temp := 0
	opt := byte('+')
	str := ""

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
			temp = temp*10 + int(s[i]-'0')
		case '+', '-':
			num = operate(num, temp, opt)
			temp = 0
			opt = s[i]
		case '(':
			str, i = cut(s, i)
			num = operate(num, calculate(str), opt)
			opt = byte('+')
			temp = 0 // 此时 temp 应该是 0 ，再次赋值，以防万一
		}
	}

	return operate(num, temp, opt)
}

func cut(s string, begin int) (string, int) {
	num := 1
	i := begin + 1
	for ; i < len(s); i++ {
		if s[i] == '(' {
			num++
		} else if s[i] == ')' {
			num--
			if num == 0 {
				break
			}
		}
	}
	return s[begin+1 : i], i
}

func operate(a, b int, opt byte) int {
	if opt == '+' {
		return a + b
	}
	return a - b
}
