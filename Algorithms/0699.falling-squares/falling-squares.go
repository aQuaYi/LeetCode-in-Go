package Problem0699

import "container/heap"

func fallingSquares(positions [][]int) []int {
	res := make([]int, 1, len(positions))
	res[0] = positions[0][1]

	sp := &segment{
		left:   positions[0][0],
		right:  positions[0][0] + positions[0][1],
		height: positions[0][1],
	}

	pq := make(PQ, 0, len(positions))
	heap.Push(&pq, sp)

	for i := 1; i < len(positions); i++ {
		sp = &segment{
			left:   positions[i][0],
			right:  positions[i][0] + positions[i][1],
			height: positions[i][1],
		}

		height := 0
		// l, r := getLeft(ss, sp), getRight(ss, sp)
		// for j := l; j <= r; j++ {
		// }

		removes := make([]*segment, 0, len(pq))

		for i := 0; i < len(pq); i++ {
			if pq[i].right <= sp.left || sp.right <= pq[i].left {
				continue
			}

			height = max(height, pq[i].height)

			if sp.left <= pq[i].left && pq[i].right <= sp.right {
				removes = append(removes, pq[i])
			}

			if pq[i].left < sp.left && sp.right < pq[i].right {
				heap.Push(&pq, &segment{
					left:   sp.right,
					right:  pq[i].right,
					height: pq[i].height,
				})
				pq[i].right = sp.left
				break
			}

			if pq[i].left < sp.left && sp.left < pq[i].right && pq[i].right <= sp.right {
				pq[i].right = sp.left
			}

			if sp.left <= pq[i].left && pq[i].left < sp.right && sp.right < pq[i].right {
				pq[i].left = sp.right
			}
		}

		for i := 0; i < len(removes); i++ {
			heap.Remove(&pq, removes[i].index)
		}

		sp.height += height
		heap.Push(&pq, sp)
		res = append(res, pq[0].height)
	}

	return res
}

// entry 是 priorityQueue 中的元素
type segment struct {
	left, right int
	height      int
	// index 是 entry 在 heap 中的索引号
	index int
}

// // 返回 a - b 后的宽度
// func minus(a, b *segment) int {

// }

// PQ implements heap.Interface and holds entries.
type PQ []*segment

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool {
	return pq[i].height > pq[j].height
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push 往 pq 中放 entry
func (pq *PQ) Push(x interface{}) {
	temp := x.(*segment)
	temp.index = len(*pq)
	*pq = append(*pq, temp)
}

// Pop 从 pq 中取出最优先的 entry
func (pq *PQ) Pop() interface{} {
	temp := (*pq)[len(*pq)-1]
	temp.index = -1 // for safety
	*pq = (*pq)[0 : len(*pq)-1]
	return temp
}

func getLeft(ss []*segment, s *segment) int {
	lo, hi := 0, len(ss)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if ss[mid].right <= s.left {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return lo
}

func getRight(ss []*segment, s *segment) int {
	lo, hi := 0, len(ss)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if s.right <= ss[mid].left {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return hi
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
