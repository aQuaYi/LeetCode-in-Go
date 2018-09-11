package problem0010

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
	two string
}

type ans struct {
	one bool
}

func Test_Problem0010(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			p: para{
				one: "aa",
				two: "a",
			},
			a: ans{
				one: false,
			},
		},
		question{
			p: para{
				one: "aa",
				two: "aa",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "aaa",
				two: "aa",
			},
			a: ans{
				one: false,
			},
		},

		question{
			p: para{
				one: "aa",
				two: "a*",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "aa",
				two: ".*",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "ab",
				two: ".*",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "aab",
				two: "c*a*b",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "aaaaaaaab",
				two: "c*a*b",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "ab",
				two: ".*c",
			},
			a: ans{
				one: false,
			},
		},
		question{
			p: para{
				one: "ab",
				two: "z*t*x*c*a*b",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "ab",
				two: "c*a*b",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "abc",
				two: ".*",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "ab",
				two: ".*b.*",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "b",
				two: ".*b.*",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "b",
				two: ".*...b",
			},
			a: ans{
				one: false,
			},
		},
		question{
			p: para{
				one: "b",
				two: ".*..*b",
			},
			a: ans{
				one: false,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, isMatch(p.one, p.two), "输入:%v", p)
	}
}
