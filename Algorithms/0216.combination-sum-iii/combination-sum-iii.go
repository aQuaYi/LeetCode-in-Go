package problem0216

func combinationSum3(k int, n int) [][]int {
	res := [][]int{}
	temp := make([]int, k+1)
	used := make([]bool, 10)

	var dfs func(int, int)
	dfs = func(idx, remain int) {
		if idx == k {
			if remain < 10 && !used[remain] {
				temp[idx] = remain
				t := make([]int, k)
				copy(t, temp[1:])
				res = append(res, t)
			}
			return
		}

		for i := temp[idx-1] + 1; i < 10; i++ {
			if remain-i  < i {
				return
			}
			used[i] = true
			temp[idx] = i
			dfs(idx+1, remain-i)
			used[i] = false
		}
	}

	dfs(1, n)

	return res
}
