package problem0871

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	target    int
	startFuel int
	stations  [][]int
	ans       int
}{

	{
		1,
		1,
		[][]int{},
		0,
	},

	{
		100,
		1,
		[][]int{{10, 100}},
		-1,
	},

	{
		100,
		10,
		[][]int{{10, 60}, {20, 30}, {30, 30}, {60, 40}},
		2,
	},

	// 可以有多个 testcase
}

func Test_minRefuelStops(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, minRefuelStops(tc.target, tc.startFuel, tc.stations), "输入:%v", tc)
	}
}

func Benchmark_minRefuelStops(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			minRefuelStops(tc.target, tc.startFuel, tc.stations)
		}
	}
}
