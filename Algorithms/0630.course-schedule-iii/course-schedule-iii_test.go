package problem0630

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	courses [][]int
	ans     int
}{

	{
		[][]int{[]int{5, 5}, []int{4, 6}, []int{2, 6}},
		2,
	},

	{
		[][]int{[]int{100, 200}, []int{200, 1300}, []int{1000, 1250}, []int{2000, 3200}},
		3,
	},

	// 可以有多个 testcase
}

func Test_scheduleCourse(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, scheduleCourse(tc.courses), "输入:%v", tc)
	}
}

func Benchmark_scheduleCourse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			scheduleCourse(tc.courses)
		}
	}
}
