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
	// h[i] > h[j] 使得 h[0] == max(h...)
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

type edgeSlice [][3]int

func (es edgeSlice) Len() int {
	return len(es)
}
func (es edgeSlice) Less(i, j int) bool {
	if es[i][0] != es[j][0] {
		return es[i][0] < es[j][0]
	}
	// 当 es[i][0] == es[j][0] 的时候
	// i,j 分别为 进入边 与 退出边
	//     则，进入边应在前
	// i,j 同为 进入边
	//     则，e[1] 高的在前
	// i,j 同为 退出边
	//     则，e[1] 低的在前
	return es[i][1]*es[i][2] < es[j][1]*es[j][2]
}
func (es edgeSlice) Swap(i, j int) {
	es[i], es[j] = es[j], es[i]
}

func getSkyline(buildings [][]int) [][]int {
	res := [][]int{}

	edges := make([][3]int, 0, 2*len(buildings))
	for _, b := range buildings {
		// e[2] == -1 为 进入边
		edges = append(edges, [3]int{b[0], b[2], -1})
		// e[2] == 1 为 退出边
		edges = append(edges, [3]int{b[1], b[2], 1})
	}
	sort.Sort(edgeSlice(edges))

	high := new(highHeap)
	heap.Init(high)
	pre := 0
	heap.Push(high, 0)

	for _, e := range edges {
		if e[2] < 0 {
			heap.Push(high, e[1])
		} else {
			i := 0
			for i < len(*high) {
				if (*high)[i] == e[1] {
					break
				}
				i++
			}
			heap.Remove(high, i)
		}
		now := (*high)[0]
		if pre != now {
			res = append(res, []int{e[0], now})
			pre = now
		}
	}

	return res
}
