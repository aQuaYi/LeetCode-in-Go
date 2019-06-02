package problem0996

import "math"

// ref: https://leetcode.com/problems/number-of-squareful-arrays/discuss/238562/C%2B%2BPython-Backtracking
func numSquarefulPerms(A []int) int {
	size := len(A)

	count := make(map[int]int, size)
	for _, a := range A {
		count[a]++
	}

	cands := make(map[int][]int, size)
	for x := range count {
		for y := range count {
			if isSquare(x + y) {
				cands[x] = append(cands[x], y)
			}
		}
	}

	res := 0
	var dfs func(int, int)
	dfs = func(x, remain int) {
		if remain == 0 {
			res++
			return
		}
		count[x]--
		for _, y := range cands[x] {
			if count[y] > 0 {
				dfs(y, remain-1)
			}
		}
		count[x]++
	}

	for x := range count {
		dfs(x, size-1)
	}

	return res
}

func isSquare(x int) bool {
	root := int(math.Sqrt(float64(x)))
	return root*root == x
}
