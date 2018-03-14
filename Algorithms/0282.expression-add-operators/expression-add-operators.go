package problem0282

func addOperators(s string, target int) []string {
	res := []string{}

	var dfs func(string, string, int, int)

	dfs = func(s, resStr string, result, prevAdd int) {
		var currStr, nextS string
		var currNum int

		// s 切分完成
		if len(s) == 0 {
			// 检查是否符合 target
			if result == target {
				res = append(res, resStr)
			}
			return
		}

		for i := 1; i <= len(s); i++ {
			currStr = s[:i]
			// 排除类似 "01" 的数
			if currStr[0] == '0' && len(currStr) > 1 {
				// 不是 continue
				// 因为出现了一次 "01" 以后，后面的 str 都只会是 "01X" 之类的。
				return
			}

			currNum = integer(currStr)
			nextS = s[i:]

			if len(resStr) == 0 {
				dfs(nextS, currStr, currNum, currNum)
			} else {
				/*
					当 运算符 opt 为 + 或 - 时，把计算式
					result = result opt currNum
					统一成
					result += prevAdd, prevAdd = 0 opt currNum
					这样做的原因是，当 opt 为 * 时，可以先利用 prevAdd 保证 乘法 运算的优先性
					result -= prevAdd
					prevAdd *= currNum
					result += prevAdd
				*/
				dfs(nextS, resStr+"+"+currStr, result+currNum, currNum)
				dfs(nextS, resStr+"-"+currStr, result-currNum, -currNum)
				dfs(nextS, resStr+"*"+currStr, result-prevAdd+prevAdd*currNum, prevAdd*currNum)
			}
		}
	}

	dfs(s, "", 0, 0)

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
