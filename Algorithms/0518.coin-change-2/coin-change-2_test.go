package problem0518

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	amount int
	coins  []int
	ans    int
}{

	{
		5,
		[]int{1, 2, 5},
		4,
	},

	{
		3,
		[]int{2},
		0,
	},

	{
		10,
		[]int{10},
		1,
	},

	{
		500,
		[]int{3, 5, 7, 8, 9, 10, 11},
		35502874,
	},

	// 可以有多个 testcase
}

func Test_myFunc(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, change(tc.amount, tc.coins), "输入:%v", tc)
	}
}

func Benchmark_myFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			change(tc.amount, tc.coins)
		}
	}
}
