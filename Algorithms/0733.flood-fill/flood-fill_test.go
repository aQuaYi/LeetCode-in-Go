package problem0733

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	image    [][]int
	sr       int
	sc       int
	newColor int
	ans      [][]int
}{

	{
		[][]int{{0, 0, 0}, {0, 1, 1}},
		1,
		1,
		1,
		[][]int{{0, 0, 0}, {0, 1, 1}},
	},

	{
		[][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}},
		1,
		1,
		2,
		[][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}},
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, floodFill(tc.image, tc.sr, tc.sc, tc.newColor), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			floodFill(tc.image, tc.sr, tc.sc, tc.newColor)
		}
	}
}
