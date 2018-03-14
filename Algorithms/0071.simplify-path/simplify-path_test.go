package problem0071

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
	path string
}

// ans 是答案
type ans struct {
	one string
}

func Test_Problem0071(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				"/home/",
			},
			ans{
				"/home",
			},
		},

		question{
			para{
				"/a/./b/../../c/",
			},
			ans{
				"/c",
			},
		},

		question{
			para{
				"/../",
			},
			ans{
				"/",
			},
		},

		question{
			para{
				"/home//foo/",
			},
			ans{
				"/home/foo",
			},
		},

		question{
			para{
				"/home/./foo/",
			},
			ans{
				"/home/foo",
			},
		},

		question{
			para{
				"/home////////////////////////foo/",
			},
			ans{
				"/home/foo",
			},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, simplifyPath(p.path), "输入:%v", p)
	}
}

func Benchmark_NewSlice(b *testing.B) {
	dir := make([]byte, 0, 3)
	dir = dir[:0]
	for i := 0; i < b.N; i++ {
		dir = []byte{}
	}
}

func Benchmark_CutSlice(b *testing.B) {
	dir := make([]byte, 0, 3)
	for i := 0; i < b.N; i++ {
		dir = dir[:0]
	}
}
