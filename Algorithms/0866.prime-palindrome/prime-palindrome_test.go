package problem0866

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	N   int
	ans int
}{

	{
		1000000000,
		10000500001,
	},

	{
		11,
		11,
	},

	{
		8,
		11,
	},

	{
		98390,
		98689,
	},

	{
		983910,
		1003001,
	},

	{
		930,
		10301,
	},

	{
		192,
		313,
	},

	{
		102,
		131,
	},

	{
		12,
		101,
	},

	{
		13,
		101,
	},

	{
		9999,
		10301,
	},

	{
		6,
		7,
	},

	{
		2,
		2,
	},

	{
		3,
		3,
	},

	{
		5,
		5,
	},

	{
		7,
		7,
	},

	// 可以有多个 testcase
}

func Test_primePalindrome(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, primePalindrome(tc.N), "输入:%v", tc)
	}
}

func Benchmark_primePalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			primePalindrome(tc.N)
		}
	}
}

func Test_isPrime(t *testing.T) {
	ast := assert.New(t)

	actual := isPrime(3)
	expected := true
	ast.Equal(expected, actual)

}
