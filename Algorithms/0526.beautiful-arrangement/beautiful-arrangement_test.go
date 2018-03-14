package problem0526

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
	N int
}

// ans 是答案
type ans struct {
	one int
}

func Test_Problem0526(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				2,
			},
			ans{
				2,
			},
		},

		question{
			para{
				3,
			},
			ans{
				3,
			},
		},

		question{
			para{
				4,
			},
			ans{
				8,
			},
		},

		question{
			para{
				5,
			},
			ans{
				10,
			},
		},

		question{
			para{
				6,
			},
			ans{
				36,
			},
		},

		question{
			para{
				7,
			},
			ans{
				41,
			},
		},

		question{
			para{
				8,
			},
			ans{
				132,
			},
		},

		question{
			para{
				9,
			},
			ans{
				250,
			},
		},

		question{
			para{
				10,
			},
			ans{
				700,
			},
		},

		question{
			para{
				11,
			},
			ans{
				750,
			},
		},

		question{
			para{
				12,
			},
			ans{
				4010,
			},
		},

		question{
			para{
				13,
			},
			ans{
				4237,
			},
		},

		question{
			para{
				14,
			},
			ans{
				10680,
			},
		},

		question{
			para{
				15,
			},
			ans{
				24679,
			},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, countArrangement(p.N), "输入:%v", p)
	}
}
