package problem0038

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
	one string
}

func Test_Problem0038(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{1},
			ans{"1"},
		},

		question{
			para{2},
			ans{"11"},
		},

		question{
			para{3},
			ans{"21"},
		},

		question{
			para{4},
			ans{"1211"},
		},

		question{
			para{5},
			ans{"111221"},
		},

		question{
			para{6},
			ans{"312211"},
		},

		question{
			para{7},
			ans{"13112221"},
		},

		question{
			para{20},
			ans{"11131221131211132221232112111312111213111213211231132132211211131221131211221321123113213221123113112221131112311332211211131221131211132211121312211231131112311211232221121321132132211331121321231231121113112221121321133112132112312321123113112221121113122113121113123112112322111213211322211312113211"},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, countAndSay(p.one), "输入:%v", p)
	}
}
