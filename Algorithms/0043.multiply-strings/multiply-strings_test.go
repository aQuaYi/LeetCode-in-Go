package problem0043

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
	one string
}

func Test_Problem0043(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{"9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999"},
			ans{"99999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999980000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001"},
		},

		question{
			para{"123", "65"},
			ans{"7995"},
		},

		question{
			para{"123", "0"},
			ans{"0"},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, multiply(p.one, p.two), "输入:%v", p)
	}
}
