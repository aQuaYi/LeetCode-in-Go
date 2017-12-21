package Problem0678

func checkValidString(s string) bool {
	// high == x 代表了，
	// 在 for 循环中 s[:i+1] 中的 '*' 全部变成 '('，最多能有 x 个 '('
	// 所以，
	// 		遇到了 '(' 或者 '*'，要 high++
	// 		遇到了 ')' 要 high--
	high := 0
	// low == x 代表了，
	// 在 for 循环中 s[:i+1] 中的 '*' 全部变成 ')'，最少能有 x 个 '('
	// 所以，本来
	// 		遇到了 '(' 要 low++
	// 		遇到了 ')' 或 '*' 要 low--
	// 但是，我们注意到
	// 当 high = -1 < 0 时，
	// 		可知就算 s[:i] 中的 '*' 全部变成了 '(' ，
	// 		仍然没有与 s[i]==')' 相匹配的 '('，
	// 		此时可以断定 s[:i+1] 是不可行的，直接返回 false
	// 当 low = -1 < 0 时，
	// 		却不能得出这样的结论
	// 		因为过多的 ')' 可能是 '*' 变过来的
	// 所以，low 的更新方式应该是
	// 		遇到了 '(' 要 low++
	// 		遇到了 ')' 或 '*' 要 low-- if low > 0
	low := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			// 遇到 '(' ，
			// low 和 high 同时增加
			high++
			low++
		case ')':
			// 遇到 ')'
			high--
			if low > 0 {
				low--
			}
			if high < 0 {
				// 此时 low == 0，因此，也可以写成
				// if high < low {
				// 表示使得 s 成立的可行解空间不存在
				// 因此，返回 false
				return false
			}
		default:
			high++
			if low > 0 {
				low--
			}
		}
	}

	return low == 0
}

// 由以上的程序过程可知
// 在导致 return false 的 high-- 之前
// 始终都有 0 <= low <= high
// 所以，[low, high] 是使得 s 成立的所有可行解的一种映射
