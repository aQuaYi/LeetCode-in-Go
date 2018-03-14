package problem0149

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	xy  [][]int
	ans int
}{

	{
		[][]int{[]int{-54, -297}, []int{-36, -222}, []int{3, -2}, []int{30, 53}, []int{-5, 1}, []int{-36, -222}, []int{0, 2}, []int{1, 3}, []int{6, -47}, []int{0, 4}, []int{2, 3}, []int{5, 0}, []int{48, 128}, []int{24, 28}, []int{0, -5}, []int{48, 128}, []int{-12, -122}, []int{-54, -297}, []int{-42, -247}, []int{-5, 0}, []int{2, 4}, []int{0, 0}, []int{54, 153}, []int{-30, -197}, []int{4, 5}, []int{4, 3}, []int{-42, -247}, []int{6, -47}, []int{-60, -322}, []int{-4, -2}, []int{-18, -147}, []int{6, -47}, []int{60, 178}, []int{30, 53}, []int{-5, 3}, []int{-42, -247}, []int{2, -2}, []int{12, -22}, []int{24, 28}, []int{0, -72}, []int{3, -4}, []int{-60, -322}, []int{48, 128}, []int{0, -72}, []int{-5, 3}, []int{5, 5}, []int{-24, -172}, []int{-48, -272}, []int{36, 78}, []int{-3, 3}},
		30,
	},

	{
		[][]int{},
		0,
	},

	{
		[][]int{
			[]int{1, 1},
			[]int{1, 1},
			[]int{1, 1},
		},
		3,
	},

	// 可以有多个 testcase
}

func Test_maxPoints(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		points := makePoints(tc.xy)
		ast.Equal(tc.ans, maxPoints(points), "输入:%v", tc)
	}
}

func Benchmark_maxPoints(b *testing.B) {
	points := makePoints(
		[][]int{[]int{-54, -297}, []int{-36, -222}, []int{3, -2}, []int{30, 53}, []int{-5, 1}, []int{-36, -222}, []int{0, 2}, []int{1, 3}, []int{6, -47}, []int{0, 4}, []int{2, 3}, []int{5, 0}, []int{48, 128}, []int{24, 28}, []int{0, -5}, []int{48, 128}, []int{-12, -122}, []int{-54, -297}, []int{-42, -247}, []int{-5, 0}, []int{2, 4}, []int{0, 0}, []int{54, 153}, []int{-30, -197}, []int{4, 5}, []int{4, 3}, []int{-42, -247}, []int{6, -47}, []int{-60, -322}, []int{-4, -2}, []int{-18, -147}, []int{6, -47}, []int{60, 178}, []int{30, 53}, []int{-5, 3}, []int{-42, -247}, []int{2, -2}, []int{12, -22}, []int{24, 28}, []int{0, -72}, []int{3, -4}, []int{-60, -322}, []int{48, 128}, []int{0, -72}, []int{-5, 3}, []int{5, 5}, []int{-24, -172}, []int{-48, -272}, []int{36, 78}, []int{-3, 3}},
	)

	for i := 0; i < b.N; i++ {
		maxPoints(points)
	}
}

func makePoints(xy [][]int) []Point {
	res := make([]Point, len(xy))
	for i := 0; i < len(xy); i++ {
		res[i] = Point{X: xy[i][0], Y: xy[i][1]}
	}

	return res
}
