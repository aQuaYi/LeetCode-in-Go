package problem0749

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	grid [][]int
	ans  int
}{

	{
		[][]int{
			{0, 1, 0, 1, 1, 1, 1, 1, 1, 0},
			{0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
			{0, 0, 1, 1, 1, 0, 0, 0, 1, 0},
			{0, 0, 0, 1, 1, 0, 0, 1, 1, 0},
			{0, 1, 0, 0, 1, 0, 1, 1, 0, 1},
			{0, 0, 0, 1, 0, 1, 0, 1, 1, 1},
			{0, 1, 0, 0, 1, 0, 0, 1, 1, 0},
			{0, 1, 0, 1, 0, 0, 0, 1, 1, 0},
			{0, 1, 1, 0, 0, 1, 1, 0, 0, 1},
			{1, 0, 1, 1, 0, 1, 0, 1, 0, 1},
		},
		38,
	},

	{
		[][]int{
			{0, 1, 0, 0, 0, 0, 0, 1},
			{0, 1, 0, 0, 0, 0, 0, 1},
			{0, 0, 0, 0, 0, 0, 0, 1},
			{0, 0, 0, 0, 0, 0, 0, 0}},
		10,
	},

	{
		[][]int{
			{1, 1, 1},
			{1, 0, 1},
			{1, 1, 1}},
		4,
	},

	{
		[][]int{
			{1, 1, 1, 0, 0, 0, 0, 0, 0},
			{1, 0, 1, 0, 1, 1, 1, 1, 1},
			{1, 1, 1, 0, 0, 0, 0, 0, 0}},
		13,
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, containVirus(tc.grid), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			containVirus(tc.grid)
		}
	}
}
