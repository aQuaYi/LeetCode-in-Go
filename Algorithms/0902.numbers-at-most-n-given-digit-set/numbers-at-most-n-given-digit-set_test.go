package problem0902

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	D   []string
	N   int
	ans int
}{

	{
		[]string{"9"},
		55,
		1,
	},

	{
		[]string{"7"},
		8,
		1,
	},

	{
		[]string{"1", "3", "5", "7"},
		1357,
		112,
	},

	{
		[]string{"1", "3", "5", "7"},
		200,
		36,
	},

	{
		[]string{"1", "3", "5", "7"},
		100,
		20,
	},

	{
		[]string{"1", "4", "9"},
		1000000000,
		29523,
	},

	// 可以有多个 testcase
}

func Test_atMostNGivenDigitSet(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, atMostNGivenDigitSet(tc.D, tc.N), "输入:%v", tc)
	}
}

func Benchmark_atMostNGivenDigitSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			atMostNGivenDigitSet(tc.D, tc.N)
		}
	}
}

func Test_zeroLead(t *testing.T) {
	type args struct {
		size   int
		length int
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		{
			"4 个数字，组成最多 2 位数的组合数",
			args{
				4,
				3,
			},
			20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zeroLead(tt.args.size, tt.args.length); got != tt.want {
				t.Errorf("empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_less(t *testing.T) {
	type args struct {
		size   int
		length int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"4 个不同的数，组合 5 位数的组合数",
			args{
				4,
				6, // 5+1
			},
			1024,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := less(tt.args.size, tt.args.length); got != tt.want {
				t.Errorf("less() = %v, want %v", got, tt.want)
			}
		})
	}
}
