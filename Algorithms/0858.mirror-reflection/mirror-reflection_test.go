package problem0858

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	p   int
	q   int
	ans int
}{

	{
		2,
		2,
		1,
	},

	{
		2,
		1,
		2,
	},

	// 可以有多个 testcase
}

func Test_mirrorReflection(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, mirrorReflection(tc.p, tc.q), "输入:%v", tc)
	}
}

func Benchmark_mirrorReflection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			mirrorReflection(tc.p, tc.q)
		}
	}
}
