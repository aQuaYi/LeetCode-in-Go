package Problem0731

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MyCanlendar(t *testing.T) {
	ast := assert.New(t)

	mc := Constructor()

	ast.True(mc.Book(10, 20))
	ast.True(mc.Book(50, 60))
	ast.True(mc.Book(10, 40))

	ast.False(mc.Book(5, 15))

	ast.True(mc.Book(5, 10))
	ast.True(mc.Book(25, 55))
}

func Test_MyCanlendar_2(t *testing.T) {
	events := [][]int{{5, 12}, {42, 50}, {4, 9}, {33, 41}, {2, 7}, {16, 25}, {7, 16}, {6, 11}, {13, 18}, {38, 43}, {49, 50}, {6, 15}, {5, 13}, {35, 42}, {19, 24}, {46, 50}, {39, 44}, {28, 36}, {28, 37}, {20, 29}, {41, 49}, {11, 19}, {41, 46}, {28, 37}, {17, 23}, {22, 31}, {4, 10}, {31, 40}, {4, 12}, {19, 26}}
	ans := []bool{true, true, true, true, false, true, false, false, true, true, true, false, false, false, true, false, false, true, false, false, false, false, false, false, false, false, false, false, false, false}

	ast := assert.New(t)

	mc := Constructor()

	for i, e := range events {
		ast.Equal(ans[i], mc.Book(e[0], e[1]), "输入是 %d, %v", i, e)
	}
}
