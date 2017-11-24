package Problem0587

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	points [][]int
	ans    [][]int
}{

	{
		[][]int{{0, 0}, {1, 0}, {1, 3}, {1, 8}, {1, 9}, {2, 0}, {2, 6}, {3, 0}, {3, 1}, {3, 4}, {3, 6}, {4, 2}, {4, 6}, {5, 7}, {5, 8}, {6, 2}, {6, 4}, {7, 7}, {7, 9}, {8, 0}, {8, 1}, {8, 3}, {8, 5}, {9, 6}},
		[][]int{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {8, 0}, {9, 6}, {7, 9}, {1, 9}},
	},

	{
		[][]int{{1, 1}, {2, 2}, {2, 0}, {2, 4}, {3, 3}, {4, 2}},
		[][]int{{2, 0}, {4, 2}, {3, 3}, {2, 4}, {1, 1}},
	},

	{
		[][]int{{1, 2}, {2, 2}, {4, 2}},
		[][]int{{1, 2}, {2, 2}, {4, 2}},
	},

	// 可以有多个 testcase
}

func Test_outerTrees(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		points := kit.Intss2Points(tc.points)
		ans := kit.Intss2Points(tc.ans)
		ast.Equal(ans, outerTrees(points), "输入:%v", tc)
	}
}

func Benchmark_outerTrees(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			points := kit.Intss2Points(tc.points)
			outerTrees(points)
		}
	}
}
