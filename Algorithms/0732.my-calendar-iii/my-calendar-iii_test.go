package problem0732

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MyCanlendar(t *testing.T) {
	events := [][]int{
		{10, 20},
		{50, 60},
		{10, 40},
		{5, 15},
		{5, 10},
		{25, 55},
	}
	ans := []int{1, 1, 2, 3, 3, 3}

	ast := assert.New(t)

	mc := Constructor()

	for i, e := range events {
		// fmt.Println(i, e)
		ast.Equal(ans[i], mc.Book(e[0], e[1]), "输入是 %d, %v", i, e)
		// fmt.Println(mc.events)
	}
}
