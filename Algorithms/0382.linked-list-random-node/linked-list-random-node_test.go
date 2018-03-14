package problem0382

import (
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

func Test_Constructor(t *testing.T) {
	ast := assert.New(t)

	ints := []int{1, 2, 3, 4, 5, 6}
	head := kit.Ints2List(ints)
	sol := Constructor(head)

	for i := 0; i < 200; i++ {
		ast.Contains(ints, sol.GetRandom(), "从 %v 中随机获取一个数", ints)
	}
}
