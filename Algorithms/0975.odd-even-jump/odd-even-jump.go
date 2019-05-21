package problem0975

import "sort"

// ref: https://leetcode.com/problems/odd-even-jump/discuss/217981/JavaC%2B%2BPython-DP-idea-Using-TreeMap-or-Stack
func oddEvenJumps(A []int) int {
	n := len(A)

	AI := make([][2]int, n)
	for i, a := range A {
		AI[i] = [2]int{a, i}
	}

	sort.Slice(AI, func(i int, j int) bool {
		if AI[i][0] == AI[j][0] {
			return AI[i][1] < AI[j][1]
		}
		return AI[i][0] < AI[j][0]
	})

	nextHigher := make([]int, n)
	stack := make([]int, 0, n)
	for _, ai := range AI {
		i := ai[1]
		for len(stack) > 0 && stack[len(stack)-1] < i {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			nextHigher[pop] = i
		}
		stack = append(stack, i)
	}

	sort.Slice(AI, func(i int, j int) bool {
		if AI[i][0] == AI[j][0] {
			return AI[i][1] < AI[j][1]
		}
		return AI[i][0] > AI[j][0]
	})

	nextLower := make([]int, n)
	stack = make([]int, 0, n)
	for _, ai := range AI {
		i := ai[1]
		for len(stack) > 0 && stack[len(stack)-1] < i {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			nextLower[pop] = i
		}
		stack = append(stack, i)
	}

	higher := make([]int, n)
	lower := make([]int, n)
	higher[n-1] = 1
	lower[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		higher[i] = lower[nextHigher[i]]
		lower[i] = higher[nextLower[i]]
	}
	return sum(higher)
}

func sum(A []int) int {
	res := 0
	for _, a := range A {
		res += a
	}
	return res
}
