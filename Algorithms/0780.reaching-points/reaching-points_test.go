package problem0780

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	sx  int
	sy  int
	tx  int
	ty  int
	ans bool
}{

	{
		3,
		3,
		12,
		9,
		true,
	},

	{
		1,
		1,
		100000001,
		2,
		true,
	},

	{
		1,
		1,
		1000000000,
		1,
		true,
	},

	{
		1,
		1,
		3,
		5,
		true,
	},

	{
		1,
		1,
		2,
		2,
		false,
	},

	{
		1,
		1,
		1,
		1,
		true,
	},

	// 可以有多个 testcase
}

func Test_reachingPoints(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, reachingPoints(tc.sx, tc.sy, tc.tx, tc.ty), "输入:%v", tc)
	}
}

func Benchmark_reachingPoints(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			reachingPoints(tc.sx, tc.sy, tc.tx, tc.ty)
		}
	}
}
