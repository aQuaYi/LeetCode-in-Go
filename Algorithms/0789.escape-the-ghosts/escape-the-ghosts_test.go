package problem0789

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	ghosts [][]int
	target []int
	ans    bool
}{

	{
		[][]int{{-1, 0}, {0, 1}, {-1, 0}, {0, 1}, {-1, 0}},
		[]int{0, 0},
		true,
	},

	{
		[][]int{{1, 0}, {0, 3}},
		[]int{0, 1},
		true,
	},

	{
		[][]int{{1, 0}},
		[]int{2, 0},
		false,
	},

	{
		[][]int{{2, 0}},
		[]int{1, 0},
		false,
	},

	// 可以有多个 testcase
}

func Test_escapeGhosts(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, escapeGhosts(tc.ghosts, tc.target), "输入:%v", tc)
	}
}

func Benchmark_escapeGhosts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			escapeGhosts(tc.ghosts, tc.target)
		}
	}
}
