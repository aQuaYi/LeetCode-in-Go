package problem0803

import "container/heap"

func hitBricks(grid [][]int, hits [][]int) []int {
	res := make([]int, len(hits))

	for idx, hit := range hits {
		i, j := hit[0], hit[1]
		grid[i][j] = 0

		for k := 0; k < 4; k++ {
			x := i + dx[k]
			y := j + dy[k]
			if isHanging(x, y, grid) {
				continue
			}

			res[idx] += drop(x, y, grid)
		}
	}

	return res
}

func isHanging(i, j int, grid [][]int) bool {
	m, n := len(grid), len(grid[0])

	if i < 0 || m <= i || j < 0 || n <= j {
		return false
	}

	isVisited := make([]bool, m*n)
	isVisited[i*n+j] = true

	pq := make(PQ, 0, m*n)
	heap.Push(&pq, &entry{
		x: i,
		y: j,
	})

	for len(pq) > 0 {
		e := heap.Pop(&pq).(*entry)
		if e.x == 0 {
			return true
		}
		for k := 0; k < 4; k++ {
			x := e.x + dx[k]
			y := e.y + dy[k]
			if 0 <= x && x < m && 0 <= y && y < n && grid[x][y] == 1 && !isVisited[x*n+y] {
				heap.Push(&pq, &entry{
					x: x,
					y: y,
				})
				isVisited[x*n+y] = true
			}
		}
	}

	return false
}

func drop(i, j int, grid [][]int) int {
	m, n := len(grid), len(grid[0])

	if i < 0 || m <= i ||
		j < 0 || n <= j ||
		grid[i][j] == 0 {
		return 0
	}

	grid[i][j] = 0
	res := 1

	for k := 0; k < 4; k++ {
		x := i + dx[k]
		y := j + dy[k]
		res += drop(x, y, grid)
	}

	return res
}

var dx = []int{-1, 0, 0, 1}
var dy = []int{0, -1, 1, 0}

// entry 是 priorityQueue 中的元素
type entry struct {
	x, y int
}

// PQ implements heap.Interface and holds entries.
type PQ []*entry

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool {
	return pq[i].x < pq[j].x
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push 往 pq 中放 entry
func (pq *PQ) Push(x interface{}) {
	temp := x.(*entry)
	*pq = append(*pq, temp)
}

// Pop 从 pq 中取出最优先的 entry
func (pq *PQ) Pop() interface{} {
	temp := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return temp
}
