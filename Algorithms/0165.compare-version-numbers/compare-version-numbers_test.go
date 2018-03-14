package problem0165

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	version1 string
	version2 string
	ans      int
}{


	{
		"1",
		"1.0.1",
		-1,
	},

	{
		"1.0.1",
		"1",
		1,
	},

	{
		"1",
		"0",
		1,
	},

	{
		"0.1",
		"0.2",
		-1,
	},

	{
		"0.1",
		"1.0",
		-1,
	},

	{
		"1.2",
		"1.1",
		1,
	},

	{
		"1.0",
		"0.1",
		1,
	},

	{
		"1.1",
		"1.1",
		0,
	},

	// 可以有多个 testcase
}

func Test_compareVersion(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, compareVersion(tc.version1, tc.version2), "输入:%v", tc)
	}
}

func Benchmark_compareVersion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			compareVersion(tc.version1, tc.version2)
		}
	}
}
