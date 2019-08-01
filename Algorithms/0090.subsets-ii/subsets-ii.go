package problem0090

import "sort"

func subsetsWithDup(A []int) [][]int {
	res := [][]int{}
	sort.Ints(A)

	var dfs func(int, []int)
	dfs = func(index int, temp []int) {
		t := make([]int, len(temp))
		copy(t, temp)
		res = append(res, t)

		for i := index; i < len(A); i++ {
			if i == index ||
				A[i] != A[i-1] { // 避免重复答案
				dfs(i+1, append(temp, A[i]))
			}
		}
	}

	temp := make([]int, 0, len(A))
	dfs(0, temp)

	return res
}
