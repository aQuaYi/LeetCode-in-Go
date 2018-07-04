package problem0850

import (
	"sort"
)

// 思路来自 https://leetcode.com/problems/rectangle-area-ii/discuss/137914/C++Python-Discretization-and-O(NlogN)
//
// 可以想象成，
// 根据 rectangles 中所有的 x 值，在平面上划 竖线
// 根据 rectangles 中所有的 y 值，在平面上划 横线
// 原先的长方形被划线分隔成更小的长方形
// 这些小长方形，要么单独存在，要么与别的小长方形**完全重合**
//
// 然后，从下往上，依次统计每一行，所有小长方形的面积，重合的小长方形只统计一次

const mod = 1e9 + 7

func rectangleArea(rectangles [][]int) int {
	// 提取 rectangles 中所有的 x 坐标值，并按照升序排列
	xs := getXs(rectangles)

	// idxs 记录了 x 在 xs 中的索引值，即
	// idxs[key]==val <==> xs[val]==key
	idxs := make(map[int]int, 2*len(xs))
	for idx, x := range xs {
		idxs[x] = idx
	}

	// labels[i]=[y,x1,x2,sig]
	// 其中 [y,[x1,x2]] 表示小长方形的一条横边
	// sig==1  为底边
	// sig==-1 为顶边
	labels := getLabels(rectangles)

	// 当 curY=j ，count[i]=5 时，意味着
	// 以 y=j,[xs[i],xs[i+1]] 为底边的小长方形，有 5 个
	count := make([]int, len(xs))

	curY, curXSum, area := 0, 0, 0

	for _, l := range labels {
		y, x1, x2, sig := l[0], l[1], l[2], l[3]
		area += (y - curY) * curXSum
		curY = y
		// 更新 curXSum
		curXSum = 0
		for i := idxs[x1]; i < idxs[x2]; i++ {
			// 更新每个小长方形的重合度
			count[i] += sig
		}
		for i := 0; i+1 < len(count); i++ {
			if count[i] > 0 {
				// curY 上
				// 所有小长方形**底边**长度相加
				curXSum += xs[i+1] - xs[i]
			}
		}
	}

	return area % mod
}

func getXs(rects [][]int) []int {
	size := len(rects)
	xs := make([]int, 0, size*2)
	xMap := make(map[int]bool, size*2)
	for _, r := range rects {
		xMap[r[0]] = true
		xMap[r[2]] = true
	}
	for k := range xMap {
		xs = append(xs, k)
	}
	sort.Ints(xs)
	return xs
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
		return labels[i][0] < labels[j][0]
	})

	return labels
}
