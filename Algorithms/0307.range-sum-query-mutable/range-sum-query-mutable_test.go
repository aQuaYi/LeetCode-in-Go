package problem0307

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var nums = []int{1, 3, 5}
var numsnil = []int{}

func Test_NumArray(t *testing.T) {
	ast := assert.New(t)
	numArray := Constructor(nums)

	ast.Equal(9, numArray.SumRange(0, 2))
	numArray.Update(1, 2)
	ast.Equal(8, numArray.SumRange(0, 2))
	ast.Equal(7, numArray.SumRange(1, 2))
	ast.Equal(1, numArray.SumRange(0, 0))

	numArray1 := Constructor(numsnil)
	ast.Nil(numArray1.root)

}

func Benchmark_NumArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numArray := Constructor(nums)
		numArray.SumRange(0, 2)
		numArray.Update(1, 2)
		numArray.SumRange(0, 2)
		numArray.SumRange(1, 2)
	}
}
