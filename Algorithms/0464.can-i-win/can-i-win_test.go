package problem0464

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	maxChoosableInteger int
	desiredTotal        int
	ans                 bool
}{

	{
		11,
		25,
		true,
	},

	{
		10,
		0,
		true,
	},

	{
		20,
		160,
		true,
	},

	{
		20,
		300,
		false,
	},

	{
		20,
		209,
		false,
	},

	{
		10,
		20,
		true,
	},

	{
		10,
		11,
		false,
	},

	// 可以有多个 testcase
}

func Test_canIWin(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, canIWin(tc.maxChoosableInteger, tc.desiredTotal), "输入:%v", tc)
	}
}

func Benchmark_canIWin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			canIWin(tc.maxChoosableInteger, tc.desiredTotal)
		}
	}
}
