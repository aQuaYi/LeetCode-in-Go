package problem0207

func canFinish(numCourses int, prerequisites [][]int) bool {
	// next[i][j] : i -> next[i]... ，i 是 next[i] 的先修课
	next := buildNext(numCourses, prerequisites)
	// pres[i] : i 的先修课程的**个数**
	pres := buildPres(next)

	return check(next, pres, numCourses)
}

func buildNext(n int, prereq [][]int) [][]int {
	next := make([][]int, n)

	for _, pair := range prereq {
		next[pair[1]] = append(next[pair[1]], pair[0])
	}

	return next
}

func buildPres(next [][]int) []int {
	pres := make([]int, len(next))

	for _, neighbours := range next {
		for _, n := range neighbours {
			pres[n]++
		}
	}

	return pres
}

func check(next [][]int, pres []int, numCourses int) bool {
	var i, j int

	// 第 i 个完成的课程的代号是 j
	for i = 0; i < numCourses; i++ {
		// 完成首先遇到的，无需先修课程的课程
		for j = 0; j < numCourses; j++ {
			if pres[j] == 0 {
				break
			}
		}

		// 每个课程都需要先修课
		// 出现了环路
		if j >= numCourses {
			return false
		}

		// 修改 pres[j] 为负数
		// 避免重修
		pres[j] = -1
		// 完成 j 课程后
		// j 的后续课程的，先修课程数量都可以 -1
		for _, c := range next[j] {
			pres[c]--
		}
	}

	return true
}
