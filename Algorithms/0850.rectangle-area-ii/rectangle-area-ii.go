package problem0850

import (
	"sort"
)

const mod = 1e9 + 7

func rectangleArea(rectangles [][]int) int {
	xs := getXs(rectangles)
	sort.Ints(xs)

	idxX := make(map[int]int, 2*len(xs))
	for k, x := range xs {
		idxX[x] = k
	}

	count := make([]int, len(xs))

	L := make([][]int, 0, 2*len(rectangles))

	for _, r := range rectangles {
		x1, y1, x2, y2 := r[0], r[1], r[2], r[3]
		L = append(L,
			[]int{y1, x1, x2, 1},
			[]int{y2, x1, x2, -1},
		)
	}

	sort.Slice(L, func(i int, j int) bool {
		eq0 := L[i][0] == L[j][0]
		eq1 := L[i][1] == L[j][1]
		eq2 := L[i][2] == L[j][2]
		if eq0 && eq1 && eq2 {
			return L[i][3] < L[j][3]
		}
		if eq0 && eq1 {
			return L[i][2] < L[j][2]
		}
		if eq0 {
			return L[i][1] < L[j][1]
		}
		return L[i][0] < L[j][0]
	})

	curY, curXSum, area := 0, 0, 0

	for _, l := range L {
		y, x1, x2, sig := l[0], l[1], l[2], l[3]
		area += (y - curY) * curXSum
		curY = y
		for i := idxX[x1]; i < idxX[x2]; i++ {
			count[i] += sig
		}
		curXSum = 0
		for i := 0; i+1 < len(count); i++ {
			if count[i] != 0 {
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

	return res
}
