package problem0636

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	n    int
	logs []string
	ans  []int
}{

	{
		2,
		[]string{"0:start:0", "0:start:2", "0:end:5", "1:start:6", "1:end:6", "0:end:7"},
		[]int{7, 1},
	},

	{
		2,
		[]string{"0:start:0", "1:start:2", "1:end:5", "0:end:6"},
		[]int{3, 4},
	},

	{
		2,
		[]string{"0:start:0", "0:end:0", "1:start:1", "1:end:1"},
		[]int{1, 1},
	},

	{
		8,
		[]string{
			"0:start:0",
			"1:start:5",
			"2:start:6",
			"3:start:9",
			"4:start:11",
			"5:start:12",
			"6:start:14",
			"7:start:15",
			"1:start:24",
			"1:end:29",
			"7:end:34",
			"6:end:37",
			"5:end:39",
			"4:end:40",
			"3:end:45",
			"0:start:49",
			"0:end:54",
			"5:start:55",
			"5:end:59",
			"4:start:63",
			"4:end:66",
			"2:start:69",
			"2:end:70",
			"2:start:74",
			"6:start:78",
			"0:start:79",
			"0:end:80",
			"6:end:85",
			"1:start:89",
			"1:end:93",
			"2:end:96",
			"2:end:100",
			"1:end:102",
			"2:start:105",
			"2:end:109",
			"0:end:114",
		},
		[]int{20, 14, 35, 7, 6, 9, 10, 14},
	},

	{
		3,
		[]string{
			"0:start:0",
			"1:start:2",
			"1:end:5",
			"0:end:6",
			"2:start:8",
			"2:start:10",
			"2:start:12",
			"2:end:14",
			"2:end:16",
			"2:end:18",
		},
		[]int{3, 4, 11},
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, exclusiveTime(tc.n, tc.logs), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			exclusiveTime(tc.n, tc.logs)
		}
	}
}
