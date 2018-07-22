package problem0710

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Solution(t *testing.T) {
	ast := assert.New(t)

	N := 100000
	blacklist := make([]int, 0, N)
	for i := 1; i < N; i++ {
		blacklist = append(blacklist, i)
	}
	s := Constructor(N, blacklist)

	for i := 0; i < 100000; i++ {
		ast.Equal(0, s.Pick())
	}
}
func Test_Solution_2(t *testing.T) {
	ast := assert.New(t)

	N := 2
	blacklist := make([]int, 0, N)
	s := Constructor(N, blacklist)

	for i := 0; i < 10; i++ {
		p := s.Pick()
		ast.True(p == 1 || p == 0, "%d", i)
	}
}

func Test_Solution_3(t *testing.T) {
	ast := assert.New(t)

	N := 10
	blacklist := []int{0, 2, 4, 6, 8}
	s := Constructor(N, blacklist)

	for i := 0; i < 10; i++ {
		ast.NotContains(blacklist, s.Pick())
	}
}
