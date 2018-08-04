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

func isWithin(rect, point []int) bool {
	x1, y1, x2, y2 := rect[0], rect[1], rect[2], rect[3]
	x, y := point[0], point[1]

	return x1 <= x && x <= x2 &&
		y1 <= y && y <= y2
}
