package problem0837

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	N   int
	K   int
	W   int
	ans float64
}{

	{
		421,
		400,
		47,
		0.71188,
	},

	{
		3,
		1,
		1,
		1.00000,
	},

	{
		10,
		1,
		10,
		1.00000,
	},

	{
		6,
		1,
		10,
		0.60000,
	},

	{
		21,
		17,
		10,
		0.73278,
	},

	// 可以有多个 testcase
}

func Test_new21Game(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.InEpsilon(tc.ans, new21Game(tc.N, tc.K, tc.W), 0.00001, "输入:%v", tc)
	}
}

func Benchmark_new21Game(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			new21Game(tc.N, tc.K, tc.W)
		}
	}
}
