package Problem0207

func canFinish(num int, prerequisites [][]int) bool {
	before := make([][]int, num)
	needPre := make([]bool, num)
	taken := make([]bool, num)

	for _, p := range prerequisites {
		// 题目说了，没有重复的 p
		before[p[0]] = append(before[p[0]], p[1])
		needPre[p[0]] = true
	}

	var dfs func(int, int) bool
	dfs = func(i, level int) bool {
		if level == num {
			// 出现了环路
			return false
		}

		if taken[i] {
			return true
		}

		if !needPre[i] {
			taken[i] = true
			return true
		}

		for _, c := range before[i] {
			if !dfs(c, level+1) {
				return false
			}
		}
		return true
	}

	for i := 0; i < num; i++ {
		if !dfs(i, 0) {
			return false
		}
	}
	return true
}
