package problem0424

func characterReplacement(s string, k int) int {
	// count 在 for 循环中，记录了
	// s[left:right+1] 中的各个字母出现的次数
	// max 保存着整个 for 循环中， count 中每次创下的新高
	var max, left int
	count := [128]int{}

	// 在 for 循环中， s[left:right+1]
	//   要么，由于 max 变大， 向右 扩张 一格
	//   要么，由于 max 不变， 向右 移动 一格
	// 但是，s[left:right+1] 的长度，始终是　characterReplacement(s[:right+1], k) 的解
	// 虽然，s[left:right+1] 不一定是那个符合题意的 subString
	for right := 0; right < len(s); right++ {
		count[s[right]]++
		max = Max(max, count[s[right]])

		// right - left + 1 - max == k 的含义是
		// 在 s[left:right+1] 中有 max 个相同的字母 X 和 k 个不同于 X 的字母
		// 通过 k 次修改后， s[left:right+1] 就全是 X 了
		// 此时 s[left:right+1] 就是满足 characterReplacement(s[:right+1], k) 的解的那个 subString

		//  right - left + 1 - max > k  的含义是
		// s[left:right+1] 中，不同于 X 的字母，多于 k 个
		// 无法通过 k 次修改，把 s[left:right+1] 中的字母全部变成 X
		// 只好把 s[left:right+1] 向右移动一格
		// 以便让 len(s[left:right+1]) 始终是 characterReplacement(s[:right+1], k) 的解
		if right-left+1-max > k {
			// 减去 s[left] 的统计值
			count[s[left]]--
			left++
		}
	}

	// len(s) == right + 1
	// 所以，最终解是 len(s) - left
	return len(s) - left

	// 其实，这个答案更直观
	// return min(len(s), max + k)
	// 使用 min 是因为有可能不用修改 k 次，s 就变成全部一样的字母了
}

// Max 返回 a 和 b 中的较大值
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
