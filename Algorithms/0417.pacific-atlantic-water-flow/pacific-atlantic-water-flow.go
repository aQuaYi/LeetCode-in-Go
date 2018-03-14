package problem0417

func pacificAtlantic(mat [][]int) [][]int {
	res := [][]int{}
	if len(mat) == 0 || len(mat[0]) == 0 {
		return res
	}

	m, n := len(mat), len(mat[0])

	// p[i][j] 表示，[i][j] 可以让水流到 Pacific 的点
	// a[i][j] 表示，[i][j] 可以让水流到 Atlantic 的点
	p, a := make([][]bool, m), make([][]bool, m)
	for i := 0; i < m; i++ {
		p[i] = make([]bool, n)
		a[i] = make([]bool, n)
	}

	// pQueue 是所有能够让水流到 Pacific 的点的队列
	// aQueue 是所有能够让水流到 Atlantic 的点的队列
	// 初始化 pQueue 和 aQueue
	pQueue := [][]int{}
	aQueue := [][]int{}
	for i := 0; i < m; i++ {
		p[i][0] = true
		pQueue = append(pQueue, []int{i, 0})
		a[i][n-1] = true
		aQueue = append(aQueue, []int{i, n - 1})
	}
	for j := 0; j < n; j++ {
		p[0][j] = true
		pQueue = append(pQueue, []int{0, j})
		a[m-1][j] = true
		aQueue = append(aQueue, []int{m - 1, j})
	}

	ds := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	bfs := func(queue [][]int, rec [][]bool) {
		for len(queue) > 0 {
			c := queue[0]
			queue = queue[1:]
			for _, d := range ds {
				i, j := c[0]+d[0], c[1]+d[1]
				if 0 <= i && i < m &&
					0 <= j && j < n &&
					!rec[i][j] &&
					mat[c[0]][c[1]] <= mat[i][j] {
					rec[i][j] = true
					queue = append(queue, []int{i, j})
				}
			}
		}
	}

	bfs(pQueue, p)
	bfs(aQueue, a)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if p[i][j] && a[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}

	return res
}
