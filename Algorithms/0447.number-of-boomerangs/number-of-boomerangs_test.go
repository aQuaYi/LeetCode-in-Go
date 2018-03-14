package problem0447

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	points [][]int
	ans    int
}{

	{
		[][]int{{1, 0}, {2, 0}},
		0,
	},

	{
		[][]int{{0, 0}, {1, 0}, {2, 0}},
		2,
	},

	// 可以有多个 testcase
}

func Test_numberOfBoomerangs(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, numberOfBoomerangs(tc.points), "输入:%v", tc)
	}
}

func Benchmark_numberOfBoomerangs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			numberOfBoomerangs(tc.points)
		}
	}
}
