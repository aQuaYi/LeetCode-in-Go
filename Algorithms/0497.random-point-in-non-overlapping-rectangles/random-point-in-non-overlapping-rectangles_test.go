package problem0497

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Solution(t *testing.T) {
	ast := assert.New(t)

	rects := [][]int{
		[]int{1, 1, 5, 5},
	}

	s := Constructor(rects)

	for i := 0; i < 10000; i++ {
		p := s.Pick()
		actual := 0
		for _, r := range rects {
			if isWithin(r, p) {
				actual++
			}
		}
		ast.Equal(1, actual)
	}
}

func Test_Solution_2(t *testing.T) {
	ast := assert.New(t)

	rects := [][]int{
		{-2, -2, -1, -1},
		{1, 0, 3, 0},
	}

	s := Constructor(rects)

	for i := 0; i < 10000; i++ {
		p := s.Pick()
		actual := 0
		for _, r := range rects {
			if isWithin(r, p) {
				actual++
			}
		}
		ast.Equal(1, actual)
	}
}

func Test_Solution_3(t *testing.T) {
	ast := assert.New(t)

	rects := [][]int{
		{0, 0, 1, 1},
		{2, 0, 3, 1},
	}

	s := Constructor(rects)

	counts := make([]int, len(rects))
	picks := 1000000

	for i := 0; i < picks; i++ {
		p := s.Pick()
		for i, r := range rects {
			if isWithin(r, p) {
				counts[i]++
			}
		}
	}

	delta := 0.001
	for i, c := range counts {
		w, h := s.rects[i][2], s.rects[i][3]
		expected := float64(w*h) / float64(s.total)
		actual := float64(c) / float64(picks)
		ast.InDelta(expected, actual, delta)
	}

}

func isWithin(rect, point []int) bool {
	x1, y1, x2, y2 := rect[0], rect[1], rect[2], rect[3]
	x, y := point[0], point[1]

	return x1 <= x && x <= x2 &&
		y1 <= y && y <= y2
}
