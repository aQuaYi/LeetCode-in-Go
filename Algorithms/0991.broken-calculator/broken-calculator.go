package problem0991

// ref: https://leetcode.com/problems/broken-calculator/discuss/234484/JavaC%2B%2BPython-Change-Y-to-X-in-1-Line
func brokenCalc(X int, Y int) int {
	res := 0
	for X < Y {
		res += Y%2 + 1
		Y = (Y + 1) / 2
	}
	return res + X - Y
}
