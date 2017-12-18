package Problem0647

func countSubstrings(s string) int {
	count := 0
	// 换一个思路
	for i := 0; i < len(s); i++ {
		// 统计所有以 i 为中心的回文
		count += extendPalindrome(s, i, i)
		// 统计所有以 i 和 i+1 为中心的回文
		count += extendPalindrome(s, i, i+1)
	}
	return count
}

func extendPalindrome(s string, left, right int) int {
	res := 0
	for left >= 0 && right < len(s) && s[left] == s[right] {
		res++
		left--
		right++
	}
	return res
}
