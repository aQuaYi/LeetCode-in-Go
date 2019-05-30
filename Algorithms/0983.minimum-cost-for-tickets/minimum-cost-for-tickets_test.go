package problem0983

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	days  []int
	costs []int
	ans   int
}{

	{
		[]int{1, 4, 6, 7, 8, 20},
		[]int{7, 2, 15},
		6,
	},

	{
		[]int{1, 4, 6, 7, 8, 20},
		[]int{2, 7, 15},
		11,
	},

	{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31},
		[]int{2, 7, 15},
		17,
	},

	// 可以有多个 testcase
}

func Test_mincostTickets(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, mincostTickets(tc.days, tc.costs), "输入:%v", tc)
	}
}

func Benchmark_mincostTickets(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			mincostTickets(tc.days, tc.costs)
		}
	}
}

func Test_min(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		{
			"a 最小",
			args{a: 0, b: 1, c: 2},
			0,
		},

		{
			"b 最小",
			args{a: 1, b: 0, c: 2},
			0,
		},

		{
			"c 最小",
			args{a: 2, b: 1, c: 0},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := min(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("min() = %v, want %v", got, tt.want)
			}
		})
	}
}
