package problem0652

import (
	"fmt"
	"sort"
	"strconv"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	pre, in []int
	ansPost [][]int
}{

	{
		[]int{1, 2, 4, 3, 2, 4, 4},
		[]int{4, 2, 1, 4, 2, 3, 4},
		[][]int{{4, 2}, {4}},
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		root := kit.PreIn2Tree(tc.pre, tc.in)
		ans := findDuplicateSubtrees(root)
		ansPost := convert(ans)
		ast.Equal(normalize(tc.ansPost), normalize(ansPost))
	}
}

func convert(roots []*TreeNode) [][]int {
	res := make([][]int, len(roots))
	for i := range res {
		res[i] = kit.Tree2Postorder(roots[i])
	}
	return res
}

func normalize(intss [][]int) string {
	ss := make([]string, len(intss))
	for _, ints := range intss {
		s := ""
		for _, n := range ints {
			s += strconv.Itoa(n)
		}
		ss = append(ss, s)
	}
	sort.Strings(ss)

	res := ""
	for _, s := range ss {
		res += s + "+"
	}

	return res
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root := kit.PreIn2Tree(tc.pre, tc.in)
			findDuplicateSubtrees(root)
		}
	}
}
