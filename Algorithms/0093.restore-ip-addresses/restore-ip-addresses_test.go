package problem0093

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
	s string
}

// ans 是答案
type ans struct {
	one []string
}

func Test_Problem0093(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				"24525511135",
			},
			ans{
				[]string{
					"245.255.11.135",
					"245.255.111.35",
				},
			},
		},

		question{
			para{
				"26526511135",
			},
			ans{
				[]string{},
			},
		},

		question{
			para{
				"000",
			},
			ans{
				[]string{},
			},
		},

		question{
			para{
				"10111",
			},
			ans{
				[]string{"1.0.1.11", "1.0.11.1", "10.1.1.1"},
			},
		},

		question{
			para{
				"111111",
			},
			ans{
				[]string{"1.1.1.111", "1.1.11.11", "1.1.111.1", "1.11.1.11", "1.11.11.1", "1.111.1.1", "11.1.1.11", "11.1.11.1", "11.11.1.1", "111.1.1.1"},
			},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, restoreIpAddresses(p.s), "输入:%v", p)
	}
}
