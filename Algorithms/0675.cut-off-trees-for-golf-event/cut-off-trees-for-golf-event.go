package Problem0675

import "container/heap"

func cutOffTree(forest [][]int) int {
	if len(forest) == 0 || len(forest[0]) == 0 {
		return -1
	}
	m, n := len(forest), len(forest[0])

	pq := make(PQ, 0, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if forest[i][j] > 1 {
				pq = append(pq, &tree{height: forest[i][j], x: i, y: j})
			}
		}
	}
	heap.Init(&pq)

	res := 0
	sx, sy := 0, 0
	for len(pq) > 0 {
		next := pq.Pop().(*tree)
		ex, ey := next.x, next.y
		steps, isAccessible := search(forest, m, n, sx, sy, ex, ey)
		if isAccessible {
			res += steps
		} else {
			return -1
		}
	}

	return res
}

var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

// 在 forest 中，
// 如果可以从 (sx,sy) 到 (ex,ey) 则返回 步数 和 true
// 如果无法到达，则返回 -1 和 false
func search(forest [][]int, m, n, sx, sy, ex, ey int) (int, bool) {
	steps, n := 0, 1
	xs := make([]int, 1, m*n)
	xs[0] = sx
	ys := make([]int, 1, m*n)
	ys[0] = sy

	for len(xs) > 0 {
		if n == 0 {
			steps++
			n = len(xs)
		}
		tx := xs[0]
		xs = xs[1:]
		ty := ys[0]
		ys = ys[1:]
		n--

		if tx == ex && ty == ey {
			return steps, true
		}

		for i := 0; i < 4; i++ {
			x := tx + dx[i]
			y := ty + dy[i]
			if 0 <= x && x < m &&
				0 <= y && y < n &&
				forest[x][y] > 0 {
				xs = append(xs, x)
				ys = append(ys, y)
			}
		}
	}

	return -1, false
}

// tree 是 priorityQueue 中的元素
type tree struct {
	height int
	x, y   int
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
