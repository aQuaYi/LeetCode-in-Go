package Problem0282

import (
	"fmt"
)

func addOperators(s string, target int) []string {
	res := []string{}

	var dfs func(string, string, int, int)

	dfs = func(s, resStr string, currRes, prevNum int) {
		var currStr, nextS string
		var currNum int
		if len(s) == 0 {
			if currRes == target {
				res = append(res, resStr)
			}
			return
		}

		for i := 1; i <= len(s); i++ {
			currStr = s[:i]
			if currStr[0] == '0' && len(currStr) > 1 {
				// 不是 continue
				return
			}

			currNum = integer(currStr)
			nextS = s[i:]

			if len(resStr) == 0 {
				dfs(nextS, currStr, currNum, currNum)
				continue
			}

			dfs(nextS, resStr+"+"+currStr, currRes+currNum, currNum)
			dfs(nextS, resStr+"-"+currStr, currRes-currNum, -currNum)
			dfs(nextS, resStr+"*"+currStr, (currRes-prevNum)+prevNum*currNum, prevNum*currNum)
		}
	}

	dfs(s, "", 0, 0)

	return res
}

func compute(ss []string) int {
	var n1, n2, n3 int
	var opt1, opt2 string
	// n1 opt1 n2 opt2 n3
	// ↓   ↓   ↓   ↓   ↓
	// 1   +   2   *   3
	n1, opt1, n2 = 0, "+", integer(ss[0])

	for i := 1; i < len(ss); i += 2 {
		fmt.Println(ss)
		opt2, n3 = ss[i], integer(ss[i+1])
		if opt2 == "*" {
			n2 = operate(n2, n3, opt2)
		} else {
			n1 = operate(n1, n2, opt1)
			opt1 = opt2
			n2 = n3
		}
	}

	return operate(n1, n2, opt1)
}

func connect(ss []string) string {
	res := ss[0]
	for i := 1; i < len(ss); i++ {
		res += ss[i]
	}
	return res
}

// 把 string 转换成 int
func integer(s string) int {
	res := int(s[0] - '0')
	for i := 1; i < len(s); i++ {
		res = res*10 + int(s[i]-'0')
	}
	return res
}

func operate(a, b int, opt string) int {
	switch opt[0] {
	case '+':
		return a + b
	case '-':
		return a - b
	default:
		return a * b
	}
}
