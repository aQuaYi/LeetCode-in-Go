package problem0792

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	S     string
	words []string
	ans   int
}{

	{

		"abcde",
		[]string{"a", "bb", "acd", "ace"},
		3,
	},

	// 可以有多个 testcase
}

func Test_numMatchingSubseq(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, numMatchingSubseq(tc.S, tc.words), "输入:%v", tc)
	}
}

func Benchmark_numMatchingSubseq(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			numMatchingSubseq(tc.S, tc.words)
		}
	}
}

func Test_isMatching(t *testing.T) {
	type args struct {
		s string
		c string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{

		{
			"包含",
			args{
				s: "abc",
				c: "abc",
			},
			true,
		},

		{
			"不包含",
			args{
				s: "abc",
				c: "abd",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatching(tt.args.s, tt.args.c); got != tt.want {
				t.Errorf("isMatching() = %v, want %v", got, tt.want)
			}
		})
	}
}
