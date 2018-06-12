package problem0841

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	rooms [][]int
	ans   bool
}{

	{
		[][]int{{1, 3}, {3, 0, 1}, {2}, {0}},
		false,
	},

	{
		[][]int{{1}, {2}, {3}, {}},
		true,
	},

	// 可以有多个 testcase
}

func Test_canVisitAllRooms(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, canVisitAllRooms(tc.rooms), "输入:%v", tc)
	}
}

func Benchmark_canVisitAllRooms(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			canVisitAllRooms(tc.rooms)
		}
	}
}
