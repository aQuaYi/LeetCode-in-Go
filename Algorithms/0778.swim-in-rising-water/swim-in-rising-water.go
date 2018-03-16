package Problem0778

import "container/heap"

func swimInWater(grid [][]int) int {
	n := len(grid)

	pq := make(PQ, 0, n*n)
	isPushed := make([]bool, n*n)

	origin := &entry{
		x:   0,
		y:   0,
		max: grid[0][0],
	}
	heap.Push(&pq, origin)
	isPushed[0] = true

	dx := [4]int{0, 0, -1, 1}
	dy := [4]int{-1, 1, 0, 0}

	for {
		// 每次把沿路经过的最高海拔值中的最低者提取出来
		// 作为优先选择的路径
		pre := heap.Pop(&pq).(*entry)
		for k := 0; k < 4; k++ {
			i := pre.x + dx[k]
			j := pre.y + dy[k]

			if i < 0 || n <= i ||
				j < 0 || n <= j ||
				isPushed[i*n+j] {
				continue
			}

			if i == n-1 && j == n-1 {
				// 找到目的地后，最后再比较一下
				// 就可以返回答案了
				return max(pre.max, grid[n-1][n-1])
			}

			next := &entry{
				x:   i,
				y:   j,
				max: max(pre.max, grid[i][j]),
			}

			// next 作为候选，放入 pq
			heap.Push(&pq, next)
			isPushed[i*n+j] = true
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// entry 是 priorityQueue 中的元素
type entry struct {
	x, y int
	max  int // 沿路经过的海拔的最高值
}

// PQ implements heap.Interface and holds entries.
type PQ []*entry

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool {
	return pq[i].max < pq[j].max
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
