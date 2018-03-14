package problem0675

import "container/heap"

func cutOffTree(forest [][]int) int {
	// 题目保证了 m,n >0
	m, n := len(forest), len(forest[0])

	// 构建砍树的优先队列
	// 根据题意，先砍矮的树
	pq := make(PQ, 0, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if forest[i][j] > 1 {
				pq = append(pq, &tree{
					height: forest[i][j],
					point:  point{x: i, y: j},
				},
				)
			}
		}
	}
	heap.Init(&pq)

	res := 0
	beg := point{x: 0, y: 0}
	for len(pq) > 0 {
		next := heap.Pop(&pq).(*tree)
		end := next.point
		// 利用 bfs 搜索从 beg 到 end 的步长
		steps, isAccessible := bfs(forest, beg, end)
		if !isAccessible {
			// 无法从 beg 到 end，直接返回 -1
			return -1
		}
		res += steps
		beg = end
	}

	return res
}

var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

// 在 forest 中，
// 如果可以从 beg 到 end 则返回 步数 和 true
// 如果无法到达，则返回 -1 和 false
func bfs(forest [][]int, beg, end point) (int, bool) {
	m, n := len(forest), len(forest[0])
	isPassed := make([]bool, m*n)

	steps, cands := 0, 1
	ps := make([]point, 1, m*n)
	ps[0] = beg

	for len(ps) > 0 {
		if cands == 0 {
			steps++
			cands = len(ps)
		}

		tp := ps[0]
		ps = ps[1:]
		isPassed[tp.x*n+tp.y] = true
		cands--

		if tp == end {
			return steps, true
		}

		for i := 0; i < 4; i++ {
			x := tp.x + dx[i]
			y := tp.y + dy[i]
			p := point{x: x, y: y}
			if 0 <= x && x < m &&
				0 <= y && y < n &&
				forest[x][y] > 0 &&
				!isPassed[x*n+y] {
				ps = append(ps, p)
			}
		}
	}

	return -1, false
}

// tree 是 priorityQueue 中的元素
type tree struct {
	height int
	point
}

type point struct {
	x, y int
}

// PQ implements heap.Interface and holds entries.
type PQ []*tree

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool {
	return pq[i].height < pq[j].height
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push 往 pq 中放 tree
func (pq *PQ) Push(x interface{}) {
	temp := x.(*tree)
	*pq = append(*pq, temp)
}

// Pop 从 pq 中取出最矮的 tree
func (pq *PQ) Pop() interface{} {
	temp := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return temp
}
