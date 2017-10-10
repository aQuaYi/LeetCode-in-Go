package Problem0282

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

// 把 string 转换成 int
func integer(s string) int {
	res := int(s[0] - '0')
	for i := 1; i < len(s); i++ {
		res = res*10 + int(s[i]-'0')
	}
	return res
}
