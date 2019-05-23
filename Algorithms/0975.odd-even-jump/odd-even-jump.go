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

	nextHigher := nextIndex(indexs)

	ascToDes(A, indexs)

	nextLower := nextIndex(indexs)

	higher, lower := make([]int, size), make([]int, size)
	higher[size-1], lower[size-1] = 1, 1
	for i := size - 2; i >= 0; i-- {
		higher[i], lower[i] = lower[nextHigher[i]], higher[nextLower[i]]
	}
	return sum(higher)
}

func nextIndex(indexs []int) []int {
	size := len(indexs)
	res := make([]int, size)
	stack := make([]int, 0, size)
	for _, j := range indexs {
		for len(stack) > 0 && stack[len(stack)-1] < j {
			pop := stack[len(stack)-1]
			res[pop] = j
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, j)
	}
	return res
}

func ascToDes(A, indexs []int) {
	i, size := 0, len(A)
	for i+1 < size {
		if A[indexs[i]] != A[indexs[i+1]] {
			i++
			continue
		}
		a, j := A[indexs[i]], i+1
		for j+1 < size && A[indexs[j+1]] == a {
			j++
		}
		reverse(indexs, i, j)
		i = j + 1
	}
	reverse(indexs, 0, size-1)
}

func reverse(A []int, i, j int) {
	for i < j {
		A[i], A[j] = A[j], A[i]
		i++
		j--
	}
}

func sum(A []int) int {
	res := 0
	for _, a := range A {
		res += a
	}
	return res
}
