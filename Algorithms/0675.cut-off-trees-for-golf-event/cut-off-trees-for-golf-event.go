package Problem0675

import "container/heap"

func cutOffTree(forest [][]int) int {
	// 题目已经保证了 m,n >=0
	m, n := len(forest), len(forest[0])

	pq := make(PQ, 0, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if forest[i][j] > 1 {
				pq = append(pq, []int{forest[i][j], i, j})
			}
		}
	}
	heap.Init(&pq)

	res := 0
	beg := []int{0, 0}
	for len(pq) > 0 {
		next := heap.Pop(&pq).([]int)
		end := next[1:]
		steps, isAccessible := search(forest, beg, end)
		if isAccessible {
			res += steps

			// fmt.Println(*next)

			beg = end
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
func search(forest [][]int, beg, end []int) (int, bool) {
	m, n := len(forest), len(forest[0])

	isPassed := make([]bool, m*n)

	steps, stepLen := 0, 1
	ps := make([][]int, 1, m*n)
	ps[0] = beg

	for len(ps) > 0 {
		if stepLen == 0 {
			steps++
			stepLen = len(ps)
		}

		tp := ps[0]
		ps = ps[1:]
		isPassed[tp[0]*n+tp[1]] = true
		stepLen--

		if tp[0] == end[0] && tp[1] == end[1] {
			return steps, true
		}

		for i := 0; i < 4; i++ {
			x := tp[0] + dx[i]
			y := tp[1] + dy[i]
			p := []int{x, y}
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
type PQ [][]int

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool {
	return pq[i][0] < pq[j][0]
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push 往 pq 中放 tree
func (pq *PQ) Push(x interface{}) {
	temp := x.([]int)
	*pq = append(*pq, temp)
}

// Pop 从 pq 中取出最矮的 tree
func (pq *PQ) Pop() interface{} {
	temp := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return temp
}
