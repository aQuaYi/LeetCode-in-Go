package Problem0218

import (
	"container/heap"
	"sort"
)

// highHeap 实现了 heap 的接口
type highHeap []int

func (h highHeap) Len() int {
	return len(h)
}
func (h highHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h highHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *highHeap) Push(x interface{}) {
	// Push 使用 *h，是因为
	// Push 增加了 h 的长度
	*h = append(*h, x.(int))
}
func (h *highHeap) Pop() interface{} {
	// Pop 使用 *h ，是因为
	// Pop 减短了 h 的长度
	res := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return res
}

type edgeSlice [][]int

func (e edgeSlice) Len() int {
	return len(e)
}
func (e edgeSlice) Less(i, j int) bool {
	if e[i][0] != e[j][0] {
		return e[i][0] < e[j][0]
	}
	return e[i][1] < e[j][1]
}
func (e edgeSlice) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func getSkyline(buildings [][]int) [][]int {
	res := [][]int{}

	edges := make([][]int, 0, 2*len(buildings))
	for _, t := range buildings {
		edges = append(edges, []int{t[0], -t[2]})
		edges = append(edges, []int{t[1], t[2]})
	}
	sort.Sort(edgeSlice(edges))

	h := new(highHeap)
	heap.Init(h)
	pre := 0
	heap.Push(h, 0)

	for i := 0; i < len(edges); i++ {
		if edges[i][1] < 0 {
			heap.Push(h, -edges[i][1])
		} else {
			idx := 0
			for idx = 0; idx < len(*h); idx++ {
				if (*h)[idx] == edges[i][1] {
					break
				}
			}
			heap.Remove(h, idx)
		}
		cur := (*h)[0]
		if pre != cur {
			res = append(res, []int{edges[i][0], cur})
			pre = cur
		}
	}

	return res
}
