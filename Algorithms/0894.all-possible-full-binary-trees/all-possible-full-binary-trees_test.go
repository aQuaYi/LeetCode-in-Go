package problem0894

import (
	"fmt"
	"sort"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"
	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	N   int
	ans [][]int
}{

	{
		2,
		[][]int{},
	},

	{
		7,
		[][]int{{0, 0, 0, kit.NULL, kit.NULL, 0, 0, kit.NULL, kit.NULL, 0, 0}, {0, 0, 0, kit.NULL, kit.NULL, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, kit.NULL, kit.NULL, kit.NULL, kit.NULL, 0, 0}, {0, 0, 0, 0, 0, kit.NULL, kit.NULL, 0, 0}},
	},

	// 可以有多个 testcase
}

func Test_allPossibleFBT(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ans := allPossibleFBT(tc.N)
		ansInts := make([][]int, 0, 20)
		for _, tree := range ans {
			ansInts = append(ansInts, kit.Tree2ints(tree))
		}

		sort.Slice(tc.ans, func(i int, j int) bool {
			for k := 0; k < len(tc.ans[i]); k++ {
				if tc.ans[i][k] != tc.ans[j][k] {
					return tc.ans[i][k] < tc.ans[j][k]
				}
			}
			return false
		})

		sort.Slice(ansInts, func(i int, j int) bool {
			for k := 0; k < len(tc.ans[i]); k++ {
				if ansInts[i][k] != ansInts[j][k] {
					return ansInts[i][k] < ansInts[j][k]
				}
			}
			return false
		})

		ast.Equal(tc.ans, ansInts, "输入:%v", tc)
	}
}

func Benchmark_allPossibleFBT(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			allPossibleFBT(tc.N)
		}
	}
}
