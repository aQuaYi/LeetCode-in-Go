package Problem0219

import (
	"testing"
	"fmt"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

// para 是参数
type para struct {
	nums []int
 k int
}

// ans 是答案
type ans struct {
	one  bool 
}

func Test_Problem0219(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
					,
			},
			ans{
					,
			},
		},
	
		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, containsNearbyDuplicate(p.one), "输入:%v", p)
	}
}
