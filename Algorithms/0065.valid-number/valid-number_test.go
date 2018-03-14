package problem0065

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
	one bool
}

func Test_Problem0065(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				"0",
			},
			ans{
				true,
			},
		},

		question{
			para{
				"  0.1  ",
			},
			ans{
				true,
			},
		},

		question{
			para{
				"abc",
			},
			ans{
				false,
			},
		},

		question{
			para{
				"1 a",
			},
			ans{
				false,
			},
		},

		question{
			para{
				"2e10",
			},
			ans{
				true,
			},
		},

		question{
			para{
				"1 1",
			},
			ans{
				false,
			},
		},

		question{
			para{
				"1e",
			},
			ans{
				false,
			},
		},

		question{
			para{
				"e2",
			},
			ans{
				false,
			},
		},
		question{
			para{
				"-5",
			},
			ans{
				true,
			},
		},

		question{
			para{
				"-0",
			},
			ans{
				true,
			},
		},

		question{
			para{
				"-0",
			},
			ans{
				true,
			},
		},

		question{
			para{
				"2e0.5",
			},
			ans{
				false,
			},
		},
		question{
			para{
				"03",
			},
			ans{
				true,
			},
		},
		question{
			para{
				"-2e-3",
			},
			ans{
				true,
			},
		},

		question{
			para{
				"    ",
			},
			ans{
				false,
			},
		},

		question{
			para{
				"0.-5",
			},
			ans{
				false,
			},
		},

		question{
			para{
				"-e3",
			},
			ans{
				false,
			},
		},

		question{
			para{
				".",
			},
			ans{
				false,
			},
		},

		question{
			para{
				"-",
			},
			ans{
				false,
			},
		},

		question{
			para{
				"3.",
			},
			ans{
				true,
			},
		},

		question{
			para{
				"46.e3",
			},
			ans{
				true,
			},
		},

		question{
			para{
				"46.5e3",
			},
			ans{
				true,
			},
		},

		question{
			para{
				"46.5e-",
			},
			ans{
				false,
			},
		},

		question{
			para{
				".e2",
			},
			ans{
				false,
			},
		},

		question{
			para{
				"..2",
			},
			ans{
				false,
			},
		},

		question{
			para{
				"1e+3",
			},
			ans{
				true,
			},
		},

		question{
			para{
				"+.8",
			},
			ans{
				true,
			},
		},

		question{
			para{
				".5",
			},
			ans{
				true,
			},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, isNumber(p.s), "输入:%v", p)
	}
}
