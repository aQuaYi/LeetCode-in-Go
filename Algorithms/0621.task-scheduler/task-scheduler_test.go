package problem0621

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
	tasks []byte
	n     int
}

// ans 是答案
type ans struct {
	one int
}

func Test_Problem0621(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				[]byte{'A', 'A', 'A', 'B', 'B', 'B'},
				2,
			},
			ans{
				8,
			},
		},

		question{
			para{
				[]byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'B', 'B', 'C', 'C', 'C'},
				1,
			},
			ans{
				12,
			},
		},

		question{
			para{
				[]byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'B', 'B', 'C', 'C', 'C'},
				0,
			},
			ans{
				12,
			},
		},

		question{
			para{
				[]byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'B', 'B', 'C', 'C', 'C'},
				100,
			},
			ans{
				506,
			},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, leastInterval(p.tasks, p.n), "输入:%v", p)
	}
}
