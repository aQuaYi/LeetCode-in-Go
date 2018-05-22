package problem0031

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
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one []int
}

var qs = []question{

	question{
		para{[]int{1, 5, 4, 3, 2}},
		ans{[]int{2, 1, 3, 4, 5}},
	},

	question{
		para{[]int{1, 5, 1}},
		ans{[]int{5, 1, 1}},
	},

	question{
		para{[]int{5, 1, 1}},
		ans{[]int{1, 1, 5}},
	},

	question{
		para{[]int{1, 1}},
		ans{[]int{1, 1}},
	},

	question{
		para{[]int{1, 2, 7, 4, 3, 1}},
		ans{[]int{1, 3, 1, 2, 4, 7}},
	},

	question{
		para{[]int{1, 2, 3}},
		ans{[]int{1, 3, 2}},
	},

	question{
		para{[]int{3, 2, 1}},
		ans{[]int{1, 2, 3}},
	},

	question{
		para{[]int{1, 1, 5}},
		ans{[]int{1, 5, 1}},
	},

	question{
		para{[]int{2, 1}},
		ans{[]int{1, 2}},
	},

	question{
		para{[]int{1}},
		ans{[]int{1}},
	},

	// 如需多个测试，可以复制上方元素。
}

func Test_Problem0031(t *testing.T) {
	ast := assert.New(t)

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)
		nextPermutation(p.one)
		ast.Equal(a.one, p.one, "输入:%v", p)
	}

}

func Benchmark_Problem0031(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			nextPermutation(q.para.one)
		}
	}
}
