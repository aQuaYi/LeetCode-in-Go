package Problem0726

func countOfAtoms(formula string) string {
	res := ""

	return res
}

func cut(formula string) (string, string) {
	i := 1
	if formula[0] == '(' {
		i = jump(formula)
	}

	for i < len(formula) &&
		!isUpper(formula[i]) &&
		formula[i] != '(' {
		i++
	}
	return formula[:i], formula[i:]
}

func isUpper(b byte) bool {
	return 'A' <= b && b <= 'Z'
}

// jump 跳过了圆括号部分
func jump(s string) int {
	p := 1
	i := 1
	for i < len(s) && p > 0 {
		if s[i] == '(' {
			p++
		} else if s[i] == ')' {
			p--
		}
		i++
	}
	return i
}
