package problem0295

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MedianFinder(t *testing.T) {
	ast := assert.New(t)

	mf := Constructor()

	for i := 1; i < 10; i++ {
		mf.AddNum(i)
		ast.Equal(float64(1+i)/2, mf.FindMedian(), "median of %v %v", mf.left.intHeap, mf.right.intHeap)
	}

	mf = Constructor()
	for i := -1; i >= -10; i-- {
		mf.AddNum(i)
		ast.Equal(float64(-1+i)/2, mf.FindMedian(), "median of %v %v", mf.left.intHeap, mf.right.intHeap)
	}

	// 没有用到 Pop 这里只为了 100% 覆盖率
	mf.left.Pop()
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
