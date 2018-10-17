package problem0062

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
	m int
	n int
}

// ans 是答案
type ans struct {
	one int
}

func Test_Problem0062(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				3,
				7,
			},
			ans{
				28,
			},
		},

		question{
			para{
				23,
				12,
			},
			ans{
				193536720,
			},
		},

		question{
			para{
				100,
				100,
			},
			ans{
				4631081169483718960,
			},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, uniquePaths(p.m, p.n), "输入:%v", p)
	}
}

func Benchmark_uniquePaths(b *testing.B) {
	for i := 1; i < b.N; i++ {
		uniquePaths(100, 100)
	}
}
