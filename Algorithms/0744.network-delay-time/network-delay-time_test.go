package Problem0744

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	times [][]int
	N     int
	K     int
	ans   int
}{

	{
		[][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}},
		4,
		2,
		2,
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, networkDelayTime(tc.times, tc.N, tc.K), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			networkDelayTime(tc.times, tc.N, tc.K)
		}
	}
}
