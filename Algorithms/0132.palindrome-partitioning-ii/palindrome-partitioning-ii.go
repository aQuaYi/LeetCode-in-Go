package Problem0132

func minCut(s string) int {
	min := len(s)
	var dfs func(int, int)
	dfs = func(i, count int) {
		if isPalindrome(s[i:]) {
			if min > count {
				min = count
			}
			return
		}

		if count > min {
			// 划分次数已经大于了已有的 min，可以提前结束了
			return
		}

		j := len(s) - 1
		for ; i < j; j-- {
			if isPalindrome(s[i:j]) {
				dfs(j, count+1)
			}
		}
	}

	dfs(0, 0)

	return min
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
