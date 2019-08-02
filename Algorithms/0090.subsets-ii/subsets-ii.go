package problem0090

import "sort"

func subsetsWithDup(A []int) [][]int {
	res := [][]int{}
	sort.Ints(A)

	var dfs func(int, []int)
	dfs = func(index int, temp []int) {
		res = append(res, temp)
		n := len(temp) + 1
		for i := index; i < len(A); i++ {
			// if 语句的含义是，
			// 在 A[index:] 中，每个数字只能附着到 temp 中一次
			// 判断方法是 A[i] != A[i-1]
			// 但是 A[index] == A[index-1] 也没有关系
			// 因为 A[index-1] 不在 A[index:] 中
			// 而且，需要执行 A[i]!=A[i-1] 时，
			// 可以肯定 i>=1，所以，不需要验证 i-1>=0
			if i == index || A[i] != A[i-1] {
				dfs(i+1, append(temp, A[i])[:n:n])
			}
		}
	}

	temp := make([]int, 0, 0)
	dfs(0, temp)

	return res
}
