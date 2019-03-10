package problem0947

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	stones [][]int
	ans    int
}{

	{[][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}},
		5,
	},

	{[][]int{{0, 0}, {0, 2}, {1, 1}, {2, 0}, {2, 2}},
		3,
	},

	{[][]int{{0, 0}},
		0,
	},

	// 可以有多个 testcase
}

func Test_removeStones(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, removeStones(tc.stones), "输入:%v", tc)
	}
}

func Benchmark_removeStones(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			removeStones(tc.stones)
		}
	}
}
