package problem0864

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	grid []string
	ans  int
}{

	{
		[]string{"@...a", ".###A", "b.BCc"},
		10,
	},

	{
		[]string{"@abcdeABCDEFf"},
		-1,
	},

	{
		[]string{"@..aA", "..B#.", "....b"},
		6,
	},

	{
		[]string{"@.a.#", "###.#", "b.A.B"},
		8,
	},

	{
		[]string{"@.....", ".####.", "....#.", ".##.#.", "BC#A#.", ".c#b#a"},
		38,
	},

	{
		[]string{"@.....", ".####.", "....#.", ".##.#.", "BA#C#.", ".a#b#c"},
		38,
	},

	{
		[]string{"@.....", ".####.", "....#.", ".##.#.", "BA#A#.", ".a#b#a"},
		28,
	},

	{
		[]string{"@.....", ".####.", "....#.", ".##.#.", "CB#A#.", ".c#b#a"},
		-1,
	},

	// 可以有多个 testcase
}

func Test_shortestPathAllKeys(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, shortestPathAllKeys(tc.grid), "输入:%v", tc)
	}
}

func Benchmark_shortestPathAllKeys(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			shortestPathAllKeys(tc.grid)
		}
	}
}
