package problem0808

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	N   int
	ans float64
}{

	{
		660295675,
		1,
	},

	{
		5000,
		1,
	},

	{
		4800,
		0.99999,
	},

	{
		20000,
		1,
	},

	{
		850,
		0.96612,
	},

	{
		50,
		0.625,
	},

	// 可以有多个 testcase
}

func Test_soupServings(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.InDelta(tc.ans, soupServings(tc.N), 0.00001, "输入:%v", tc)
	}
}

func Benchmark_soupServings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			soupServings(tc.N)
		}
	}
}
