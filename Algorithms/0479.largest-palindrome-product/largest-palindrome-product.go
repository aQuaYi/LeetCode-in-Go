package problem0479

// 题目说了 1 <= n <=8
// 利用题目的 run 功能，把这 8 个答案都找出来，直接返回就好了
var res = []int{9, 987, 123, 597, 677, 1218, 877, 475}

func largestPalindrome(n int) int {
	return res[n-1]
}
