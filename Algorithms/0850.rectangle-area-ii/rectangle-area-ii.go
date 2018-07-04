package problem0850

import (
	"sort"
)

// 思路来自 https://leetcode.com/problems/rectangle-area-ii/discuss/137914/C++Python-Discretization-and-O(NlogN)

const mod = 1e9 + 7

func rectangleArea(rectangles [][]int) int {
	// 提取 rectangles 中所有的 x 坐标值，并按照升序排列
	xs := getXs(rectangles)

	idxs := make(map[int]int, 2*len(xs))
	for idx, x := range xs {
		idxs[x] = idx
	}

	count := make([]int, len(xs))

	labels := getLabels(rectangles)

	curY, curXSum, area := 0, 0, 0

	for _, l := range labels {
		y, x1, x2, sig := l[0], l[1], l[2], l[3]
		area += (y - curY) * curXSum
		curY = y
		for i := idxs[x1]; i < idxs[x2]; i++ {
			count[i] += sig
		}
		curXSum = 0
		for i := 0; i+1 < len(count); i++ {
			if count[i] > 0 {
				curXSum += xs[i+1] - xs[i]
			}
		}
	}

	return area % mod
}

func getXs(rects [][]int) []int {
	size := len(rects)
	res := make([]int, 1, size*2)
	xMap := make(map[int]bool, size*2)
	for _, r := range rects {
		xMap[r[0]] = true
		xMap[r[2]] = true
	}
	for k := range xMap {
		res = append(res, k)
	}
	sort.Ints(res)
	return res
}

func getLabels(rects [][]int) [][]int {
	labels := make([][]int, 0, 2*len(rects))
	for _, r := range rects {
		x1, y1, x2, y2 := r[0], r[1], r[2], r[3]
		labels = append(labels,
			[]int{y1, x1, x2, 1},
			[]int{y2, x1, x2, -1},
		)
	}

	// 对 labels 进行排序
	sort.Slice(labels, func(i int, j int) bool {
		eq0 := labels[i][0] == labels[j][0]
		eq1 := labels[i][1] == labels[j][1]
		eq2 := labels[i][2] == labels[j][2]
		if eq0 && eq1 && eq2 {
			return labels[i][3] < labels[j][3]
		}
		if eq0 && eq1 {
			return labels[i][2] < labels[j][2]
		}
		if eq0 {
			return labels[i][1] < labels[j][1]
		}
		return labels[i][0] < labels[j][0]
	})

	return labels
}
