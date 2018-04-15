package problem0799

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	poured     int
	queryRow   int
	queryGlass int
	ans        float64
}{

	{
		2000,
		99,
		40,
		0,
	},

	{
		2,
		0,
		0,
		1.0,
	},

	{
		1,
		1,
		1,
		0.0,
	},

	{
		2,
		1,
		1,
		0.5,
	},

	// 可以有多个 testcase
}

func Test_champagneTower(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, champagneTower(tc.poured, tc.queryRow, tc.queryGlass), "输入:%v", tc)
	}
}

func Benchmark_champagneTower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			champagneTower(tc.poured, tc.queryRow, tc.queryGlass)
		}
	}
}
