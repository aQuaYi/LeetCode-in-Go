package Problem0587

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type Point = kit.Point

func outerTrees(points []Point) []Point {
	size := len(points)
	if size < 4 {
		return points
	}

	res := []Point{}
isUsed := make([]bool, size)

	cx := points[0]
	idx := 0
	for i := 1; i < size; i++ {
		cx,idx = cloest2X(cx, points[i],idx,i)
	}

res = append(res, cx)
isUsed[idx] = true



	return res
}

// 返回更靠近 X 轴的点
func cloest2X(a, b Point,ia,ib int) (Point,int) {
	A, B := a.Y*b.X, b.Y*a.X
	if A < B ||
		(A == B && a.X < b.X) {
		return a,ia
	}
	return b,ib
}
