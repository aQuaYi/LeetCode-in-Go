package Problem0282

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	num    string
	target int
	ans    []string
}{

	{
		"123",
		6,
		[]string{"1+2+3", "1*2*3"},
	},

	{
		"232",
		8,
		[]string{"2*3+2", "2+3*2"},
	},

	{
		"105",
		5,
		[]string{"1*0+5", "10-5"},
	},

	{
		"00",
		0,
		[]string{"0+0", "0-0", "0*0"},
	},

	{
		"3456237490",
		9191,
		[]string{},
	},

	// 可以有多个 testcase
}

func Test_addOperators(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, addOperators(tc.num, tc.target), "输入:%v", tc)
	}
}

func Benchmark_addOperators(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			addOperators(tc.num, tc.target)
		}
	}
}
