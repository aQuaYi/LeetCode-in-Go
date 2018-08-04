package problem0478

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RandPoint(t *testing.T) {
	ast := assert.New(t)
	r, a, b := 3., 2., 1.
	s := Constructor(r, a, b)
	for i := 0; i < 100; i++ {
		p := s.RandPoint()
		x, y := p[0], p[1]
		actual := math.Sqrt((x-a)*(x-a) + (y-b)*(y-b))
		ast.True(actual <= r)
	}
}

func Test_RandPoint_2(t *testing.T) {
	ast := assert.New(t)
	r, a, b := 0.01, -73839.1, -3289891.
	s := Constructor(r, a, b)
	for i := 0; i < 100; i++ {
		p := s.RandPoint()
		x, y := p[0], p[1]
		actual := math.Sqrt((x-a)*(x-a) + (y-b)*(y-b))
		ast.True(actual <= r)
	}
}
