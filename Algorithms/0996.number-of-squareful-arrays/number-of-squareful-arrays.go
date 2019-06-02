package problem0996

import "math"

func numSquarefulPerms(A []int) int {
	size := len(A)

	pairs := make([][]int, size)
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if isSquare(A[i] + A[j]) {
				pairs[i] = append(pairs[i], j)
				pairs[j] = append(pairs[j], i)
			}
		}
	}

	isSeen := make(map[[12]int]bool, size)
	res := 0
	for i := 0; i < size; i++ {
		perm := [12]int{}
		perm[size-1] = A[i]
		dfs(pairs, A, perm, [12]bool{}, isSeen, i, size-1, &res)
	}
	return res
}

func isSquare(x int) bool {
	root := int(math.Sqrt(float64(x)))
	return root*root == x
}

func dfs(pairs [][]int, A []int, perm [12]int, isVisited [12]bool, isSeen map[[12]int]bool, i, count int, res *int) {
	if count == 0 {
		if isSeen[perm] {
			return
		}
		isSeen[perm] = true
		(*res)++
		return
	}

	isVisited[i] = true

	for _, j := range pairs[i] {
		if isVisited[j] {
			continue
		}
		perm[count] = A[j]
		dfs(pairs, A, perm, isVisited, isSeen, j, count-1, res)
	}

	isVisited[i] = false
}
