package problem0091

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
	one int
}

func Test_Problem0091(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				"27",
			},
			ans{
				1,
			},
		},
		question{
			para{
				"",
			},
			ans{
				0,
			},
		},

		question{
			para{
				"0",
			},
			ans{
				0,
			},
		},

		question{
			para{
				"1",
			},
			ans{
				1,
			},
		},

		question{
			para{
				"203",
			},
			ans{
				1,
			},
		},

		question{
			para{
				"119",
			},
			ans{
				3,
			},
		},
		question{
			para{
				"2011",
			},
			ans{
				2,
			},
		},

		question{
			para{
				"1192",
			},
			ans{
				3,
			},
		},

		question{
			para{
				"1102",
			},
			ans{
				1,
			},
		},

		question{
			para{
				"1100",
			},
			ans{
				0,
			},
		},

		question{
			para{
				"1190",
			},
			ans{
				0,
			},
		},

		question{
			para{
				"1128",
			},
			ans{
				3,
			},
		},

		question{
			para{
				"1",
			},
			ans{
				1,
			},
		},

		question{
			para{
				"12",
			},
			ans{
				2,
			},
		},

		question{
			para{
				"121",
			},
			ans{
				3,
			},
		},
		question{
			para{
				"1212",
			},
			ans{
				5,
			},
		},

		question{
			para{
				"12121",
			},
			ans{
				8,
			},
		},

		question{
			para{
				"121212",
			},
			ans{
				13,
			},
		},

		question{
			para{
				"1212121",
			},
			ans{
				21,
			},
		},
		question{
			para{
				"12121212",
			},
			ans{
				34,
			},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, numDecodings(p.s), "输入:%v", p)
	}
}
