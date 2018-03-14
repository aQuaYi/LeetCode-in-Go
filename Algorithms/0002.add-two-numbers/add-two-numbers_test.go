package problem0002

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one *ListNode
	two *ListNode
}

type ans struct {
	one *ListNode
}

type question struct {
	p para
	a ans
}

func makeListNode(is []int) *ListNode {
	if len(is) == 0 {
		return nil
	}

	res := &ListNode{
		Val: is[0],
	}
	temp := res

	for i := 1; i < len(is); i++ {
		temp.Next = &ListNode{Val: is[i]}
		temp = temp.Next
	}

	return res
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: makeListNode([]int{2, 4, 3}),
				two: makeListNode([]int{5, 6, 4}),
			},
			a: ans{
				one: makeListNode([]int{7, 0, 8}),
			},
		},
		question{
			p: para{
				one: makeListNode([]int{9, 8, 7, 6, 5}),
				two: makeListNode([]int{1, 1, 2, 3, 4}),
			},
			a: ans{
				one: makeListNode([]int{0, 0, 0, 0, 0, 1}),
			},
		},
		question{
			p: para{
				one: makeListNode([]int{0}),
				two: makeListNode([]int{5, 6, 4}),
			},
			a: ans{
				one: makeListNode([]int{5, 6, 4}),
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, addTwoNumbers(p.one, p.two), "输入:%v", p)
	}
}
