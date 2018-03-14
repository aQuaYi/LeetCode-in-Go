package problem0097

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
type para struct {
	s1 string
	s2 string
	s3 string
}

// ans 是答案
type ans struct {
	one bool
}

func Test_Problem0097(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				"aabcc",
				"dbbca",
				"aadbbcbcac",
			},
			ans{
				true,
			},
		},

		question{
			para{
				"aabc",
				"dbbca",
				"aadbbbaccc",
			},
			ans{
				false,
			},
		},

		question{
			para{
				"aabcc",
				"dbbca",
				"aadbbbaccc",
			},
			ans{
				false,
			},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, isInterleave(p.s1, p.s2, p.s3), "输入:%v", p)
	}
}
