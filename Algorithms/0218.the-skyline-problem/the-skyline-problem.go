package problem0218

import (
	"container/heap"
	"sort"
)

func getSkyline(buildings [][]int) [][]int {
	points := [][]int{}

	highs := new(highHeap)
	heap.Init(highs)
	// pre 是 天际线 最左边的高度
	pre := 0
	heap.Push(highs, pre)

	edges := make([][3]int, 0, 2*len(buildings))
	for _, b := range buildings {
		// e[2] == -1 为 进入边
		edges = append(edges, [3]int{b[0], b[2], -1})
		// e[2] == 1 为 退出边
		edges = append(edges, [3]int{b[1], b[2], 1})
	}
	sort.Sort(edgeSlice(edges))

	// 从左往右，依次扫描 edges
	for _, e := range edges {
		if e[2] < 0 {
			// 遇到 进入边
			// 添加 e[1] 到 high
			heap.Push(highs, e[1])
		} else {
			// 遇到 退出边
			i := 0
			for i < len(*highs) {
				if (*highs)[i] == e[1] {
					break
				}
				i++
			}
			// 从 high 中删除 e[1]
			heap.Remove(highs, i)
		}
		// now 是当前的 high 中的最大值
		now := (*highs)[0]
		if pre != now {
			// 天际线的高度发生变化
			// 需要添加 key point
			points = append(points, []int{e[0], now})
			pre = now
		}
	}

	// 结题的关键是监控 最高线 的变化
	// 出现新的最高线时，就可以添加关键点了
	return points
}

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
