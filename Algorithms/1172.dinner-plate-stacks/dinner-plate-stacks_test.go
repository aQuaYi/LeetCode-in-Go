package problem1172

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_1(t *testing.T) {
	a := assert.New(t)
	//
	d := Constructor(2)
	//
	for i := 1; i <= 5; i++ {
		d.Push(i)
	}
	//
	a.Equal(2, d.PopAtStack(0))
	//
	d.Push(20)
	d.Push(21)
	a.Equal(20, d.PopAtStack(0))
	a.Equal(21, d.PopAtStack(2))
	//
	a.Equal(5, d.Pop())
	a.Equal(4, d.Pop())
	a.Equal(3, d.Pop())
	a.Equal(1, d.Pop())
	//
	a.Equal(-1, d.Pop())
}

func Test_2(t *testing.T) {
	a := assert.New(t)
	//
	d := Constructor(1)
	//
	for i := 1; i <= 3; i++ {
		d.Push(i)
	}
	//
	a.Equal(2, d.PopAtStack(1))
	//
	//
	a.Equal(3, d.Pop())
	a.Equal(1, d.Pop())
}

func Test_3(t *testing.T) {
	a := assert.New(t)
	//
	d := Constructor(1)
	//
	for i := 1; i <= 2; i++ {
		d.Push(i)
	}
	//
	a.Equal(2, d.PopAtStack(1))
	a.Equal(-1, d.PopAtStack(1))
	a.Equal(-1, d.PopAtStack(8))
	//
	a.Equal(1, d.Pop())
	//
	for i := 1; i <= 2; i++ {
		d.Push(i)
	}
	//
	a.Equal(2, d.Pop())
	a.Equal(1, d.Pop())
}
