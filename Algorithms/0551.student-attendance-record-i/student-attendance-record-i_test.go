package problem0551

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s   string
	ans bool
}{

	{
		"PPALLP",
		true,
	},

	{
		"PPALLL",
		false,
	},

	// 可以有多个 testcase
}

func Test_checkRecord(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, checkRecord(tc.s), "输入:%v", tc)
	}
}

func Benchmark_checkRecord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			checkRecord(tc.s)
		}
	}
}
