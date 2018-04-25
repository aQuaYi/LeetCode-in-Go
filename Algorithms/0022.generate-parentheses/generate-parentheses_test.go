package problem0022

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

// para 是参数
// one 代表第一个参数
type para struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one []string
}

func Test_Problem0022(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{2},
			ans{[]string{
				"(())",
				"()()",
			}},
		},

		question{
			para{3},
			ans{[]string{
				"((()))",
				"(()())",
				"(())()",
				"()(())",
				"()()()",
			}},
		},

		question{
			para{4},
			ans{[]string{
				"(((())))",
				"((()()))",
				"((())())",
				"((()))()",
				"(()(()))",
				"(()()())",
				"(()())()",
				"(())(())",
				"(())()()",
				"()((()))",
				"()(()())",
				"()(())()",
				"()()(())",
				"()()()()",
			}},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, generateParenthesis(p.one), "输入:%v", p)
	}
}

func Benchmark_generateParenthesis(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateParenthesis(10)
	}
}
