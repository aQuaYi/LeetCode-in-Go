package Problem0391

import (
	"sort"
)

func isRectangleCover(rectangles [][]int) bool {
	var bs blps

	for _, r := range rectangles {
		bs = append(bs, makeBlps(r)...)
	}

	sort.Sort(bs)

	return !isOverlap(bs) && !hasGap(bs)
}

func hasGap(bs blps) bool {
	i := 1
	for i < len(bs) && bs[i][0] == bs[i-1][0] {
		i++
	}

	if i == len(bs) {
		return false
	}

	h := i
	if len(bs)%h != 0 {
		return true
	}

	for i < len(bs) {
		if bs[i][0] != bs[0][0]+i/h {
			return true
		}
		i++
	}

	return false
}

func isOverlap(bs blps) bool {
	var i int
	for i = 1; i < len(bs); i++ {
		if isEqual(bs[i-1], bs[i]) {
			return true
		}
	}
	return false
}

func makeBlps(r []int) blps {
	bs := make(blps, (r[2]-r[0])*(r[3]-r[1]))

	var i, j, k int
	for i = r[0]; i < r[2]; i++ {
		for j = r[1]; j < r[3]; j++ {
			bs[k] = []int{i, j}
			k++
		}
	}

	return bs
}

// blp is bottom-left point
type blp []int

func isEqual(a, b blp) bool {
	return a[0] == b[0] && a[1] == b[1]
}

type blps [][]int

func (b blps) Len() int {
	return len(b)
}

func (b blps) Less(i, j int) bool {
	if b[i][0] == b[j][0] {
		return b[i][1] < b[j][1]
	}
	return b[i][0] < b[j][0]
}

func (b blps) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
