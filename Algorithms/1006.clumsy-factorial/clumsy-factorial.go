package problem1006

var answers = []int{0, 1, 2, 6, 7}

func clumsy(N int) int {
	if N <= 4 {
		return answers[N]
	}
	a := N*(N-1)/(N-2) + (N - 3)
	b := (N - 4) * max(1, (N-5)) / max(1, (N-6))
	return a - 2*b + clumsy(N-4)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// go to look O(1) solution
// https://leetcode.com/problems/clumsy-factorial/discuss/252279/You-never-think-of-this-amazing-O(1)-solution
