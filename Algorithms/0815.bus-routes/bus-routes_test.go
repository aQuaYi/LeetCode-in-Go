package problem0815

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	routes [][]int
	S      int
	T      int
	ans    int
}{

	{
		[][]int{{1, 2, 7}, {3, 6, 7}},
		1,
		6,
		2,
	},

	// 可以有多个 testcase
}

func Test_numBusesToDestination(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, numBusesToDestination(tc.routes, tc.S, tc.T), "输入:%v", tc)
	}
}

func Benchmark_numBusesToDestination(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			numBusesToDestination(tc.routes, tc.S, tc.T)
		}
	}
}
