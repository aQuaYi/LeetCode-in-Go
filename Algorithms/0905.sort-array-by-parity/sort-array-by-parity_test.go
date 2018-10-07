package problem0905

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A []int
}{

	{
		[]int{3, 1, 2, 4},
	},

	// 可以有多个 testcase
}

func split(a []int) (evens, odds []int) {
	size := len(a)
	var i int
	for i = 0; i < size; i++ {
		if a[i]%2 == 1 {
			break
		}
	}
	return a[:i], a[i:]
}

func isAllEven(a []int) bool {
	for _, n := range a {
		if n%2 == 1 {
			return false
		}
	}
	return true
}

func isAllOdd(a []int) bool {
	for _, n := range a {
		if n%2 == 0 {
			return false
		}
	}
	return true
}

func Test_sortArrayByParity(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ans := sortArrayByParity(tc.A)
		evens, odds := split(ans)
		ast.True(isAllEven(evens))
		ast.True(isAllOdd(odds))
	}
}

func Benchmark_sortArrayByParity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			sortArrayByParity(tc.A)
		}
	}
}
