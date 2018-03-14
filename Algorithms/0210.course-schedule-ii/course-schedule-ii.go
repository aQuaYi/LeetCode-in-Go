package problem0210

func findOrder(numCourses int, prerequisites [][]int) []int {
	n, p := build(numCourses, prerequisites)
	return search(n, p)
}

func build(num int, requires [][]int) (next [][]int, pre []int) {
	// next[i][j] : i -> next[i]... ，i 是 next[i] 的先修课
	next = make([][]int, num)
	// pres[i] : i 的先修课程的**个数**
	pre = make([]int, num)

	for _, r := range requires {
		next[r[1]] = append(next[r[1]], r[0])
		pre[r[0]]++
	}

	return
}

func search(next [][]int, pre []int) []int {
	n := len(pre)
	res := make([]int, n)
	var i, j int
	// 第 i 个完成的课程的代号是 j
	for i = 0; i < n; i++ {
		// 完成首先遇到的，先修课程为 0 的课程
		for j = 0; j < n; j++ {
			if pre[j] == 0 {
				break
			}
		}
		// 每个课程都需要先修课
		// 出现了环路
		if j == n {
			return nil
		}

		// 修改 pres[j] 为负数
		// 避免重修
		pre[j] = -1

		// 完成 j 课程后
		// j 的后续课程的，先修课程数量都可以 -1
		for _, c := range next[j] {
			pre[c]--
		}

		// 把课程代号放入答案
		res[i] = j
	}

	return res
}
