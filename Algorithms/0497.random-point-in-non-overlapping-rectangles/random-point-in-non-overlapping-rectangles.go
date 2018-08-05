package problem0497

import (
	"math/rand"
	"sort"
)

// Solution object will be instantiated and called as such:
// obj := Constructor(rects);
// param_1 := obj.Pick();
type Solution struct {
	total  int   // 所有可能离散的点的个数
	counts []int // count[i] 记录了前 i 个矩形中所有点的个数
	rects  [][]int
}

// Constructor 创建 Solution
func Constructor(rects [][]int) Solution {
	size := len(rects)

	total := 0
	counts := make([]int, size)
	tmp := make([][]int, size)

	for i, r := range rects {
		x1, y1, x2, y2 := r[0], r[1], r[2], r[3]
		w, h := x2-x1+1, y2-y1+1 // +1 是因为 [a,b] 之间一共有 b-a+1 个离散的点
		total += w * h
		counts[i] = total
		tmp[i] = []int{x1, y1, w, h} // 修改 rect 的表达方式
	}

	return Solution{
		counts: counts,
		rects:  tmp,
		total:  total,
	}
}

// Pick 返回正方形内的点
func (s *Solution) Pick() []int {
	cand := rand.Intn(s.total) + 1
	// 根据 cand 找到需要返回那个矩形下的点
	i := sort.SearchInts(s.counts, cand)

	x1, y1, w, h := s.rects[i][0], s.rects[i][1], s.rects[i][2], s.rects[i][3]

	x := x1 + rand.Intn(w)
	y := y1 + rand.Intn(h)

	return []int{x, y}
}
