package problem0223

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   int
	B   int
	C   int
	D   int
	E   int
	F   int
	G   int
	H   int
	ans int
}{

	{
		-2,
		-2,
		2,
		2,
		-2,
		-2,
		2,
		2,
		16,
	},

	{
		-3,
		0,
		3,
		4,
		0,
		-1,
		9,
		2,
		45,
	},

	// 可以有多个 testcase
}

func Test_computeArea(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, computeArea(tc.A, tc.B, tc.C, tc.D, tc.E, tc.F, tc.G, tc.H), "输入:%v", tc)
	}
}

func Benchmark_computeArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			computeArea(tc.A, tc.B, tc.C, tc.D, tc.E, tc.F, tc.G, tc.H)
		}
	}
}
