package problem1031

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	L   int
	M   int
	ans int
}{

	{
		[]int{4, 5, 14, 16, 16, 20, 7, 13, 8, 15},
		3,
		5,
		109,
	},

	{
		[]int{0, 6, 5, 2, 2, 5, 1, 9, 4},
		1,
		2,
		20,
	},

	{
		[]int{3, 8, 1, 3, 2, 1, 8, 9, 0},
		3,
		2,
		29,
	},

	{
		[]int{2, 1, 5, 6, 0, 9, 5, 0, 3, 8},
		4,
		3,
		31,
	},

	// 可以有多个 testcase
}

func Test_maxSumTwoNoOverlap(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, maxSumTwoNoOverlap(tc.A, tc.L, tc.M), "输入:%v", tc)
	}
}

func Benchmark_maxSumTwoNoOverlap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			maxSumTwoNoOverlap(tc.A, tc.L, tc.M)
		}
	}
}

func Test_sums(t *testing.T) {
	type args struct {
		A []int
		l int
	}
	tests := []struct {
		name string
		args args
		want []*entry
	}{
		{
			"test",
			args{
				A: []int{0, 1, 2, 3},
				l: 3,
			},
			[]*entry{
				&entry{6, 1, 3},
				&entry{3, 0, 2},
			},
		},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sums(tt.args.A, tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sums() = %v, want %v", got, tt.want)
			}
		})
	}
}
