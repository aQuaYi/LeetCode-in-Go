package problem0299

import "strconv"

func getHint(secret string, guess string) string {
	var bull, cow int
	// countS[n] > 0 表示 数字 n 在 secret 前面出现过，且没有被匹配为 bull
	// countG[n] > 0 表示 数字 n 在 guess  前面出现过，且没有被匹配为 bull
	var countS, countG [10]int

	size := len(secret)
	for i := 0; i < size; i++ {
		ns := int(secret[i] - '0')
		ng := int(guess[i] - '0')

		if ns == ng {
			bull++
		} else {
			// 检查 ns 能否在 guess[:i] 中匹配为 cow
			if countG[ns] > 0 {
				cow++
				countG[ns]--
			} else {
				countS[ns]++
			}

			// 检查 ng 能否在 secret[:i] 中匹配为 cow
			if countS[ng] > 0 {
				cow++
				countS[ng]--
			} else {
				countG[ng]++
			}
		}
	}

	return strconv.Itoa(bull) + "A" + strconv.Itoa(cow) + "B"
}
