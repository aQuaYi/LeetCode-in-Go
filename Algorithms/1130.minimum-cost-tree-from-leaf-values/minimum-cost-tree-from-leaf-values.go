package problem1130

import "math"

// ref: https://leetcode.com/problems/minimum-cost-tree-from-leaf-values/discuss/339959/One-Pass-O(N)-Time-and-Space
func mctFromLeafValues(A []int) int {
	n := len(A)

	stack, top, pop := make([]int, n+1), 0, 0
	stack[top] = math.MaxInt64

	res := 0
	for _, a := range A {
		for stack[top] <= a {
			pop, top = stack[top], top-1
			res += pop * min(stack[top], a)
		}
		stack[top+1], top = a, top+1
	}

	for top >= 2 {
		pop, top = stack[top], top-1
		res += pop * stack[top]
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
