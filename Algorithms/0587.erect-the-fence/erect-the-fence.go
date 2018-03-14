package problem0587

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type Point = kit.Point

func outerTrees(points []Point) []Point {
	n := len(points)
	if n < 4 {
		return points
	}

	// 选择最左边的点，作为起始点
	// 如果最左边的点，不止一个，选择其中最下面的那个
	first := 0
	for i := 1; i < n; i++ {
		if points[i].X < points[first].X||
		(points[i].X == points[first].X && points[i].Y< points[first].Y) {
			first = i
		}
	}

	res := append([]Point{}, points[first])

	cur := first
	for {
		// 寻找以 cur 为圆心，顺时针方向角度最大的点
		next := (cur + 1) % n
		for i := 0; i < n; i++ {
			if i == cur {
				continue
			}
			cross := crossProductLength(points[cur], points[next], points[i])
			if cross < 0 ||
				(cross == 0 && isFurther(points[cur], points[next], points[i])) {
				next = i
			}
		}

		// 把与 cur→next 共线的点都放入 res
		for i := 0; i < n; i++ {
			if i == cur || i == first {
				// 这是已经放入过的点
				continue
			}
			if crossProductLength(points[cur], points[next], points[i]) == 0 {
				// next 也会被放入
				res = append(res, points[i])
			}
		}

		cur = next
		// cur == first 说明围栏已经封闭了
		if cur == first {
			break
		}
	}

	return res
}

// 返回 oa×ob 的长度
// 当返回值 <0 时，说明以 o 点圆心，b 点在 a 点的顺时针方向
// 当返回值 =0 时，说明以 o 点，b 点和 a 点在同一条直线长
func crossProductLength(o, a, b Point) int {
	return (a.X-o.X)*(b.Y-o.Y) - (a.Y-o.Y)*(b.X-o.X)
}

func isFurther(o, a, b Point) bool {
	return distance(o, a) < distance(o, b)
}

func distance(a, b Point) int {
	return (a.X-b.X)*(a.X-b.X) + (a.Y-b.Y)*(a.Y-b.Y)
}
