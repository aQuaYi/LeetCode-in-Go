package Problem0295

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MedianFinder(t *testing.T) {
	ast := assert.New(t)

	mf := Constructor()

	mf.AddNum(1)
	// [1]
	mf.AddNum(2)
	// [1,2]
	ast.Equal(1.5, mf.FindMedian(), "median of %v", mf.nums)
	mf.AddNum(3)
	// [1,2,3]
	ast.Equal(2.0, mf.FindMedian(), "median of %v", mf.nums)
}

func Benchmark_MedianFinder(b *testing.B) {
	size := 100
	mf := Constructor()

	for i := 0; i < b.N; i++ {
		for j := 0; j < size; j++ {
			mf.AddNum(rand.Int())
			mf.FindMedian()
		}
	}
}
