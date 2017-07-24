package Problem0084

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	p para
	a ans
}

type para struct {
	one string
}

type ans struct {
	one string
}

func Test_Problem0084(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			p: para{""},
			a: ans{""},
		},
		
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, p.one, "输入:%v", p)
	}
}
