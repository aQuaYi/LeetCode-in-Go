package problem0554

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	wall [][]int
	ans  int
}{

	{
		[][]int{{1, 2, 2, 1},
			{3, 1, 2},
			{1, 3, 2},
			{2, 4},
			{3, 1, 2},
			{1, 3, 1, 1}},
		2,
	},

	// 可以有多个 testcase
}

func Test_leastBricks(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, leastBricks(tc.wall), "输入:%v", tc)
	}
}

func Benchmark_leastBricks(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			leastBricks(tc.wall)
		}
	}
}
