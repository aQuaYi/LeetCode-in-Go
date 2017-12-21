package Problem0678

func checkValidString(s string) bool {
	// low  记录了 s 中无配对的 '(' 的下限
	// high 记录了 s 中无配对的 '(' 的上限
	low, high := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			// 遇到新的无配对的 '(' ，
			// low 和 high 同时增加
			low++
			high++
		} else if s[i] == ')' {
			// 新出现了 ')'
			// 无配对 '(' 的上限会下降
			high--
			if high < 0 {
				return false
			}
			//
			if low > 0 {
				low--
			}
		} else {
			if low > 0 {
				low--
			}
			high++
		}
	}

	return low == 0
}
