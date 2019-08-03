package problem1125

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	reqs   []string
	people [][]string
	ans    []int
}{

	{
		[]string{"java", "nodejs", "reactjs"},
		[][]string{{"java"}, {"nodejs"}, {"nodejs", "reactjs"}},
		[]int{0, 2},
	},

	{
		[]string{"algorithms", "math", "java", "reactjs", "csharp", "aws"},
		[][]string{{"algorithms", "math", "java"}, {"algorithms", "math", "reactjs"}, {"java", "csharp", "aws"}, {"reactjs", "csharp"}, {"csharp", "math"}, {"aws", "java"}},
		[]int{1, 2},
	},

	// 可以有多个 testcase
}

func Test_smallestSufficientTeam(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, smallestSufficientTeam(tc.reqs, tc.people), "输入:%v", tc)
	}
}

func Benchmark_smallestSufficientTeam(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			smallestSufficientTeam(tc.reqs, tc.people)
		}
	}
}
