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

	for i := 0; i < 100; i++ {
		ast.NotContains(blacklist, s.Pick())
	}
}
