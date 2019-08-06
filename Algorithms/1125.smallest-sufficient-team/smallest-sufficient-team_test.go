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
		[]string{"rxocajdeo", "mymflnlegtrnr", "xbyqqh", "jarzoaupviqkx", "sfgstfsqlqmfjv", "lpnpelpwy", "pyqrhjamdrias"},
		[][]string{{"pyqrhjamdrias"}, {"rxocajdeo", "pyqrhjamdrias"}, {}, {"rxocajdeo"}, {"mymflnlegtrnr", "jarzoaupviqkx"}, {}, {"rxocajdeo"}, {}, {"sfgstfsqlqmfjv"}, {"rxocajdeo"}, {"pyqrhjamdrias"}, {}, {}, {}, {"lpnpelpwy"}, {}, {"rxocajdeo", "mymflnlegtrnr"}, {"rxocajdeo"}, {"xbyqqh", "pyqrhjamdrias"}, {"mymflnlegtrnr"}},
		[]int{3, 4, 8, 14, 18},
	},

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
		skillMap := make(map[string]bool, len(tc.reqs))
		indexs := smallestSufficientTeam(tc.reqs, tc.people)
		for _, i := range indexs {
			for _, s := range tc.people[i] {
				skillMap[s] = true
			}
		}
		for _, r := range tc.reqs {
			_, ok := skillMap[r]
			ast.True(ok, "组合中应该含有 %s 技能", r)
			delete(skillMap, r)
		}
		ast.Equal(0, len(skillMap))
	}
}

func Benchmark_smallestSufficientTeam(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			smallestSufficientTeam(tc.reqs, tc.people)
		}
	}
}
