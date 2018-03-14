package problem0488

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	board string
	hand  string
	ans   int
}{

	{
		// TODO: 这个 test case 官方程序会返回 -1
		// 实际上正确答案是 2
		// RRWWRRBBRR -> RRWWRRBBR[W]R -> RRWWRRBB[B]RWR -> RRWWRRRWR -> RRWWWR -> RRR -> <empty>
		// 这也是我暂时没有做这个题目的原因

		"RRWWRRBBRR",
		"WB",
		-1,
	},

	{
		"WRRBBW",
		"RB",
		-1,
	},

	{
		"WWRRBBWW",
		"WRBRW",
		2,
	},

	{
		"G",
		"GGGGG",
		2,
	},

	{
		"RBYYBBRRB",
		"YRBGB",
		3,
	},

	// 可以有多个 testcase
}

func Test_findMinStep(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findMinStep(tc.board, tc.hand), "输入:%v", tc)
	}
}

func Benchmark_findMinStep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findMinStep(tc.board, tc.hand)
		}
	}
}
