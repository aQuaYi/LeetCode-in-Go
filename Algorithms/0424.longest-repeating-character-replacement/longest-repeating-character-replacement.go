package Problem0424

func characterReplacement(s string, k int) int {
	// count 在 for 循环中，记录了
	// s[left:right+1] 中的各个字母出现的次数
	// max 保存着整个 for 循环中， count 中的最大值
	var max, left int
	count := [128]int{}

	for right := 0; right < len(s); right++ {
		count[s[right]]++
		max = Max(max, count[s[right]])
		if max+k < right-left+1 {
			count[s[left]]--
			left++
		}
	}

	return max + k
}

// Max 返回 a 和 b 中的较大值
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
