package problem0234

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	ans  bool
}{

	{[]int{1, 2, 3, 2, 1}, true},
	{[]int{1, 3, 2, 1}, false},

	// 可以有多个 testcase
}

func Test_isPalindrome(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		head := kit.Ints2List(tc.nums)
		ast.Equal(tc.ans, isPalindrome(head), "输入:%v", tc)
	}
}

func Benchmark_isPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			head := kit.Ints2List(tc.nums)
			isPalindrome(head)
		}
	}
}
