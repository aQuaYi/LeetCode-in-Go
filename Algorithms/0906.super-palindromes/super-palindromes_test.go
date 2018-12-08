package problem0906

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	L, R string
	ans  int
}{

	{
		"1020762146323",
		"142246798855636",
		17,
	},

	{
		"398904669",
		"13479046850",
		6,
	},

	{
		"1",
		"2",
		1,
	},

	{
		"4",
		"1000",
		4,
	},

	// 可以有多个 testcase
}

func Test_superpalindromesInRange(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, superpalindromesInRange(tc.L, tc.R), "输入:%v", tc)
	}
}

func Benchmark_myFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			superpalindromesInRange(tc.L, tc.R)
		}
	}
}
