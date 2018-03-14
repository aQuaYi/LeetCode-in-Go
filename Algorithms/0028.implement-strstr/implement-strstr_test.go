package problem0028

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
	two string
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one int
}

func Test_Problem0028(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{"", ""},
			ans{0},
		},

		question{
			para{"abcd", "bc"},
			ans{1},
		},

		question{
			para{"abcde", "c"},
			ans{2},
		},

		question{
			para{"abcde", "f"},
			ans{-1},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, strStr(p.one, p.two), "输入:%v", p)
	}
}
