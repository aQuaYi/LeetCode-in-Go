package problem0321

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums1 []int
	nums2 []int
	k     int
	ans   []int
}{

	{
		[]int{6, 3, 9, 0, 5, 6},
		[]int{2, 2, 5, 2, 1, 4, 4, 5, 7, 8, 9, 3, 1, 6, 9, 7, 0},
		23,
		[]int{6, 3, 9, 2, 2, 5, 2, 1, 4, 4, 5, 7, 8, 9, 3, 1, 6, 9, 7, 0, 5, 6, 0},
	},

	{
		[]int{1, 6, 5, 4, 7, 3, 9, 5, 3, 7, 8, 4, 1, 1, 4},
		[]int{4, 3, 1, 3, 5, 9},
		21,
		[]int{4, 3, 1, 6, 5, 4, 7, 3, 9, 5, 3, 7, 8, 4, 1, 3, 5, 9, 1, 1, 4},
	},

	{
		[]int{6, 7},
		[]int{6, 0, 4},
		5,
		[]int{6, 7, 6, 0, 4},
	},

	{
		[]int{1, 1, 1},
		[]int{9, 9, 9},
		2,
		[]int{9, 9},
	},

	{
		[]int{1, 1, 1},
		[]int{1, 1, 1},
		5,
		[]int{1, 1, 1, 1, 1},
	},

	{
		[]int{2, 8, 0, 4, 5, 1, 4, 8, 9, 9, 0, 8, 2, 9},
		[]int{5, 9, 6, 6, 4, 1, 0, 7},
		22,
		[]int{5, 9, 6, 6, 4, 2, 8, 1, 0, 7, 0, 4, 5, 1, 4, 8, 9, 9, 0, 8, 2, 9},
	},

	{
		[]int{3, 4, 6, 5},
		[]int{9, 1, 2, 5, 8, 3},
		5,
		[]int{9, 8, 6, 5, 3},
	},

	{
		[]int{3, 9},
		[]int{8, 9},
		3,
		[]int{9, 8, 9},
	},

	// 可以有多个 testcase
}

func Test_maxNumber(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, maxNumber(tc.nums1, tc.nums2, tc.k), "输入:%v", tc)
	}
}

func Benchmark_maxNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			maxNumber(tc.nums1, tc.nums2, tc.k)
		}
	}
}

func Test_choose(t *testing.T) {
	type args struct {
		A []int
		k int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{

		{
			"6选3",
			args{
				[]int{8, 9, 1, 7, 5, 2},
				3,
			},
			[]int{9, 7, 5},
		},

		{
			"5选3",
			args{
				[]int{5, 2, 7, 1, 9},
				3,
			},
			[]int{7, 1, 9},
		},

		//
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := choose(tt.args.A, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("choose() = %v, want %v", got, tt.want)
			}
		})
	}
}
