package problem0818

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	target int
	ans    int
}{

	{
		5,
		7,
	},

	{
		11,
		10,
	},

	{
		8,
		6,
	},

	{
		9,
		8,
	},

	{
		10,
		7,
	},

	{
		2,
		4,
	},

	{
		10000,
		45,
	},

	{
		15,
		4,
	},

	{
		6,
		5,
	},

	{
		999,
		20,
	},

	{
		100,
		19,
	},

	{
		1000,
		23,
	},

	{
		3,
		2,
	},

	{
		4,
		5,
	},

	{
		2,
		4,
	},

	{
		1,
		1,
	},

	// 可以有多个 testcase
}

func Test_racecar(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, racecar(tc.target), "输入:%v", tc)
	}
}

func Benchmark_racecar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			racecar(tc.target)
		}
	}
}
