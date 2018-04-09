package problem0787

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	n       int
	flights [][]int
	src     int
	dst     int
	K       int
	ans     int
}{

	{
		3,
		[][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}},
		0,
		2,
		1,
		200,
	},

	{
		3,
		[][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}},
		0,
		2,
		0,
		500,
	},

	// 可以有多个 testcase
}

func Test_findCheapestPrice(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findCheapestPrice(tc.n, tc.flights, tc.src, tc.dst, tc.K), "输入:%v", tc)
	}
}

func Benchmark_findCheapestPrice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findCheapestPrice(tc.n, tc.flights, tc.src, tc.dst, tc.K)
		}
	}
}
