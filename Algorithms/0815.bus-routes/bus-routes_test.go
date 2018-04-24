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
		[][]int{{1, 9, 12, 20, 23, 24, 35, 38}, {10, 21, 24, 31, 32, 34, 37, 38, 43}, {10, 19, 28, 37}, {8}, {14, 19}, {11, 17, 23, 31, 41, 43, 44}, {21, 26, 29, 33}, {5, 11, 33, 41}, {4, 5, 8, 9, 24, 44}},
		37,
		28,
		1,
	},

	{
		[][]int{{1, 2, 7}, {3, 6, 7}},
		1,
		7,
		1,
	},

	{
		[][]int{{1, 2, 7}, {3, 6, 7}},
		1,
		6,
		2,
	},

	{
		[][]int{{1, 2, 7}, {3, 6, 7}},
		1,
		1,
		0,
	},

	{
		[][]int{{1, 2, 7}, {3, 6, 7}},
		1,
		5,
		-1,
	},

	{
		[][]int{{1, 2, 7}, {3, 6, 7}, {4, 5}},
		1,
		5,
		-1,
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
