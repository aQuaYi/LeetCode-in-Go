package problem0407

import (
	"container/heap"
)

func trapRainWater(hs [][]int) int {
	if len(hs) < 3 || len(hs[0]) < 3 {
		return 0
	}
	m, n := len(hs), len(hs[0])

	pq := make(priorityQueue, 0, m*2+n*2)

	isVisited := make([][]bool, m)
	for i := range isVisited {
		isVisited[i] = make([]bool, n)
	}

	// 把四周的格子先放入 pq
	for i := 0; i < m; i++ {
		isVisited[i][0] = true
		isVisited[i][n-1] = true
		pq = append(pq, cell{row: i, col: 0, height: hs[i][0]})
		pq = append(pq, cell{row: i, col: n - 1, height: hs[i][n-1]})
	}
	for j := 0; j < n; j++ {
		isVisited[0][j] = true
		isVisited[m-1][j] = true
		pq = append(pq, cell{row: 0, col: j, height: hs[0][j]})
		pq = append(pq, cell{row: m - 1, col: j, height: hs[m-1][j]})
	}
	// 放入后，再初始化 pq，更节约时间
	heap.Init(&pq)

	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	vol := 0

	// 此刻，pq 中的 cell 实际上把整个区域围了起来
	// 从围墙 pq 中，挑选最矮的木板 c
	// 统计 c 周围的 cell 能装多少水
	// 然后把 c 周围的 cell 更新高度后，放入 pq，变成新的木板
	// 不断重复上述过程，围墙不断变小，直到统计完成所有的 cell
	for len(pq) > 0 {
		// 最矮的木板是 c
		c := heap.Pop(&pq).(cell)
		// 依次检查 c 周围的 4 个方向的 cell
		for _, d := range dirs {
			i := c.row + d[0]
			j := c.col + d[1]

			if 0 <= i && i < m && 0 <= j && j < n && !isVisited[i][j] {
				isVisited[i][j] = true
				// 统计容量
				vol += max(0, c.height-hs[i][j])
				// 给 pq 添加新的木板
				heap.Push(&pq, cell{row: i, col: j, height: max(hs[i][j], c.height)})
			}
		}
	}

	return vol
}

type cell struct {
	row, col, height int
}

type priorityQueue []cell

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].height < pq[j].height
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push 往 pq 中放 item
func (pq *priorityQueue) Push(x interface{}) {
	item := x.(cell)
	*pq = append(*pq, item)
}

// Pop 从 pq 中取出最优先的 item
func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
