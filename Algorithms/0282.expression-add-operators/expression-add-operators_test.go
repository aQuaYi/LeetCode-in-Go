package problem0282

import (
	"fmt"
	"sort"
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
		"1210",
		120,
		[]string{"12*10"},
	},

	{
		"100100",
		10000,
		[]string{"100*100"},
	},

	{
		"011",
		0,
		[]string{"0*1*1", "0*11", "0+1-1", "0-1+1"},
	},

	{
		"17",
		17,
		[]string{"17"},
	},

	{
		"",
		5,
		[]string{},
	},

	{
		"5",
		5,
		[]string{"5"},
	},

	{
		"5",
		3,
		[]string{},
	},

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

	{
		"000",
		0,
		[]string{"0*0*0", "0*0+0", "0*0-0", "0+0*0", "0+0+0", "0+0-0", "0-0*0", "0-0+0", "0-0-0"},
	},

	// 可以有多个 testcase
}

func Test_addOperators(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		sort.Strings(tc.ans)
		ans := addOperators(tc.num, tc.target)
		sort.Strings(ans)
		ast.Equal(tc.ans, ans, "输入:%v", tc)
	}
}

func Benchmark_addOperators(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			addOperators(tc.num, tc.target)
		}
	}
}
