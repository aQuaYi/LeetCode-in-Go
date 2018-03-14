package problem0609

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	path []string
	ans  [][]string
}{

	{
		[]string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)", "root 4.txt(efgh)"},
		[][]string{{"root/a/2.txt", "root/c/d/4.txt", "root/4.txt"}, {"root/a/1.txt", "root/c/3.txt"}},
	},

	// 可以有多个 testcase
}

func Test_fcName(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ans := findDuplicate(tc.path)
		ast.Equal(normal(tc.ans), normal(ans), "输入:%v", tc)
	}
}

func normal(sss [][]string) string {
	ss := make([]string, len(sss))

	for i := 0; i < len(sss); i++ {
		sort.Strings(sss[i])
		ss[i] = strings.Join(sss[i], "☆")
	}

	sort.Strings(ss)

	return strings.Join(ss, "★★")
}

func Benchmark_fcName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findDuplicate(tc.path)
		}
	}
}
