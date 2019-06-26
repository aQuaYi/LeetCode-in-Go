package problem1029

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	costs [][]int
	ans   int
}{

	{
		[][]int{{259, 770}, {448, 54}, {926, 667}, {184, 139}, {840, 118}, {577, 469}},
		1859,
	},

	{
		[][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}},
		110,
	},

	// 可以有多个 testcase
}

func Test_twoCitySchedCost(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, twoCitySchedCost(tc.costs), "输入:%v", tc)
	}
}

func Benchmark_twoCitySchedCost(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			twoCitySchedCost(tc.costs)
		}
	}
}
