package problem0003

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one string
}

type ans struct {
	one int
}

type question struct {
	p para
	a ans
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: "abcabcbb",
			},
			a: ans{
				one: 3,
			},
		},
		question{
			p: para{
				one: "bbbbbbbb",
			},
			a: ans{
				one: 1,
			},
		},
		question{
			p: para{
				one: "pwwkew",
			},
			a: ans{
				one: 3,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, lengthOfLongestSubstring(p.one), "输入:%v", p)
	}
}
