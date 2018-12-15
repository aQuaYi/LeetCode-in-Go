package problem0020

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
	one string
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one bool
}

func Test_Problem0020(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{"()[]{}"},
			ans{true},
		},
		question{
			para{"(]"},
			ans{false},
		},
		question{
			para{"({[]})"},
			ans{true},
		},
		question{
			para{"(){[({[]})]}"},
			ans{true},
		},
		question{
			para{"((([[[{{{"},
			ans{false},
		},
		question{
			para{"(())]]"},
			ans{false},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, isValid(p.one), "输入:%v", p)
	}
}

func Benchmark_isValid(b *testing.B) {
	for i := 1; i < b.N; i++ {
		isValid("{{{{{[[[[[((((()))))]]]]]}}}}}")
	}
}
