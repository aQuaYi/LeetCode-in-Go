package problem0975

import "sort"

// ref: https://leetcode.com/problems/odd-even-jump/discuss/217981/JavaC%2B%2BPython-DP-idea-Using-TreeMap-or-Stack
func oddEvenJumps(A []int) int {
	size := len(A)

	indexs := make([]int, size)
	for i := range indexs {
		indexs[i] = i
	}

	sort.Slice(indexs, func(i int, j int) bool {
		if A[indexs[i]] == A[indexs[j]] {
			return indexs[i] < indexs[j]
		}
		return A[indexs[i]] < A[indexs[j]]
	})

	nextHigher := make([]int, size)
	stack := make([]int, 0, size)
	for _, j := range indexs {
		for len(stack) > 0 && stack[len(stack)-1] < j {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			nextHigher[pop] = j
		}
		stack = append(stack, j)
	}

	sort.Slice(indexs, func(i int, j int) bool {
		if A[indexs[i]] == A[indexs[j]] {
			return indexs[i] < indexs[j]
		}
		return A[indexs[i]] > A[indexs[j]]
	})

	nextLower := make([]int, size)
	stack = make([]int, 0, size)
	for _, j := range indexs {
		for len(stack) > 0 && stack[len(stack)-1] < j {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			nextLower[pop] = j
		}
		stack = append(stack, j)
	}

	higher, lower := make([]int, size), make([]int, size)
	higher[size-1], lower[size-1] = 1, 1
	for i := size - 2; i >= 0; i-- {
		higher[i], lower[i] = lower[nextHigher[i]], higher[nextLower[i]]
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
