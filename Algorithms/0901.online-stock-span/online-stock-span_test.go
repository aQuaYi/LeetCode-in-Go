package problem0901

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_StockSpanner(t *testing.T) {
	ast := assert.New(t)
	//
	nexts := []int{100, 80, 60, 70, 60, 75, 85}
	expecteds := []int{1, 1, 1, 2, 1, 4, 6}
	//
	s := Constructor()
	for i, n := range nexts {
		expected := expecteds[i]
		actual := s.Next(n)
		ast.Equal(expected, actual)
	}
}

func Test_StockSpanner_2(t *testing.T) {
	ast := assert.New(t)
	//
	nexts := []int{29, 91, 62, 76, 51}
	expecteds := []int{1, 2, 1, 2, 1}
	//
	s := Constructor()
	for i, n := range nexts {
		expected := expecteds[i]
		actual := s.Next(n)
		ast.Equal(expected, actual)
	}
}
