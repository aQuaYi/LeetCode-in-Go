package Problem0488

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
		"RRWWRRBBRR",
		"WB",
		2,
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
