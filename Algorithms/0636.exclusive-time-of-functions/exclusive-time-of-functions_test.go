package Problem0636

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
	2
	["0:start:0","0:start:2","0:end:5","1:start:6","1:end:6","0:end:7"]
	[]int{7,1},
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

	{
		2,
		[]string{"0:start:0", "1:start:2", "1:end:5", "0:end:6"},
		[]int{3, 4},
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
