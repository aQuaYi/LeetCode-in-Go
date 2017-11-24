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

	return dfs(points) 
}

// 当 p 在 points 围起来的范围内时，返回 true
func isIn(ps []Point, i, j, k, l int) (bool, int) {
	switch {
	case isInTrangle(ps, i, j, k, l):
		return true, l
	case isInTrangle(ps, l, i, j, k):
		return true, k
	case isInTrangle(ps, k, l, i, j):
		return true, j
	case isInTrangle(ps, j, k, l, i):
		return true, i
	}
	return false, -1
}

func isInTrangle(ps []Point, i, j, k, l int) bool {
	a, b, c, p := ps[i], ps[j], ps[k], ps[l]
	std := area(a, b, c)
	pa := area(a, b, p) + area(b, c, p) + area(c, a, p)
	return pa == std
}

// 返回由 a，b，c 三点组成的三角形的面积的 2 倍
func area(a, b, c Point) int {
	return b.X*c.Y + a.X*b.Y + a.Y*c.X - b.X*a.Y - c.X*b.Y - c.Y*a.X
}

func dfs(ps []Point) []Point {
	size := len(ps)

	for i := 0; i < size-3; i++ {
		for j := i + 1; j < size-2; j++ {
			for k := j + 1; k < size-1; k++ {
				for l := k + 1; l < size; l++ {
					in, idx := isIn(ps, i, j, k, l)
					if in {
						temp := make([]Point, size-1)
						copy(temp[:idx], ps)
						copy(temp[idx:], ps[idx+1:])
						return dfs(temp)
					}
				}
			}
		}
	}
	return ps
}
