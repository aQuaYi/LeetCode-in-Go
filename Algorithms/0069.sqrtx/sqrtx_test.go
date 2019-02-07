package problem0069

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

// para 是参数
type para struct {
	x int
}

// ans 是答案
type ans struct {
	one int
}

func Test_Problem0069(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				15,
			},
			ans{
				3,
			},
		},

		question{
			para{
				3,
			},
			ans{
				1,
			},
		},

		question{
			para{
				9,
			},
			ans{
				3,
			},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, mySqrt(p.x), "输入:%v", p)
	}
}

func Benchmark_mySqrt(b *testing.B) {
	n := 10000
	for i := 1; i < b.N; i++ {
		_ = mySqrt(n)
	}
}

func Benchmark_math_sqrt(b *testing.B) {
	n := 10000
	for i := 1; i < b.N; i++ {
		_ = int(math.Sqrt(float64(n)))
	}
}
