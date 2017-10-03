package Problem0218

import (
	"container/heap"
	"sort"
)

type ints []int

func (s ints) Less(i, j int) bool {
	return s[i] > s[j]
}
func (s ints) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ints) Len() int {
	return len(s)
}
func (s *ints) Push(x interface{}) {
	// Push 使用 *s，是因为
	// Push 增加了 s 的长度
	*s = append(*s, x.(int))
}
func (s *ints) Pop() interface{} {
	// Pop 使用 *s，是因为
	// Pop 减短了 s 的长度
	x := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	return x
}

type bdHeap [][]int

func (h bdHeap) Len() int { return len(h) }
func (h bdHeap) Less(i, j int) bool {
	if h[i][0] != h[j][0] {
		return h[i][0] < h[j][0]
	}
	return h[i][1] < h[j][1]
}

func (h bdHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func getSkyline(buildings [][]int) [][]int {
	res := [][]int{}

	edges := make([][]int, 0, 2*len(buildings))
	for _, t := range buildings {
		edges = append(edges, []int{t[0], -t[2]})
		edges = append(edges, []int{t[1], t[2]})
	}
	sort.Sort(bdHeap(edges))

	h := new(ints)
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
