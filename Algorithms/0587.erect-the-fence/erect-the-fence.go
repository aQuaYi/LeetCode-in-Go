package Problem0587

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
	"math"
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
		cx, idx = cloest2X(cx, points[i], idx, i)
	}

	res = append(res, cx)
	isUsed[idx] = true

	var dfs func(Point, vector, int)
	dfs = func(o Point, b vector, index int) {
		max := -2.
		idx := -1
		for i, p := range points {
			if i == index {
				continue
			}
			tmp := angle(makeVector(o, p), b)
			if max < tmp ||
				(max == tmp && distance(p, o) < distance(points[idx], o)) {
				max = tmp
				idx = i
			}
		}

		if isUsed[idx] {
			return
		}

		isUsed[idx] = true
		dfs(points[idx], makeVector(o, points[idx]), idx)
	}

	dfs(cx, makeVector(Point{X: 0, Y: 0}, cx), idx)

	return res
}

// 返回更靠近 X 轴的点
func cloest2X(a, b Point, ia, ib int) (Point, int) {
	A, B := a.Y*b.X, b.Y*a.X
	if A < B ||
		(A == B && a.X < b.X) {
		return a, ia
	}
	return b, ib
}

type vector struct {
	x, y float64
}

// 返回 a→b 的向量
func makeVector(a, b Point) vector {
	return vector{
		x: float64(b.X - a.X),
		y: float64(b.Y - a.Y),
	}
}

// 返回两个向量的夹角
func angle(a, b vector) float64 {
	ab := a.x*a.x + a.y*a.y
	aNorm := math.Sqrt(a.x*a.x + a.y*a.y)
	bNorm := math.Sqrt(b.x*b.x + b.y*b.y)
	return ab / (aNorm * bNorm)
}

func distance(a, b Point) int {
	x := a.X - b.X
	y := a.Y - b.Y
	return x*x + y*y
}
