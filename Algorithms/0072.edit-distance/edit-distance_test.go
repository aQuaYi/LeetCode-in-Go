package problem0072

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
	from, to string
}

// ans 是答案
type ans struct {
	one int
}

func Test_Problem0072(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				"esa",
				"eat",
			},
			ans{
				2,
			},
		},

		question{
			para{
				"sea",
				"eat",
			},
			ans{
				2,
			},
		},

		question{
			para{
				"horse",
				"ros",
			},
			ans{
				3,
			},
		},

		question{
			para{
				"",
				"",
			},
			ans{
				0,
			},
		},

		question{
			para{
				"gaod",
				"bad",
			},
			ans{
				2,
			},
		},

		question{
			para{
				"gd",
				"bad",
			},
			ans{
				2,
			},
		},

		question{
			para{
				"good",
				"bad",
			},
			ans{
				3,
			},
		},

		question{
			para{
				"hello",
				"world",
			},
			ans{
				4,
			},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, minDistance(p.from, p.to), "输入:%v", p)
	}
}
