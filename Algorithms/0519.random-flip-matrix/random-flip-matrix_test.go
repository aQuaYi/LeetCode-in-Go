package problem0519

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Solution(t *testing.T) {
	ast := assert.New(t)

	rows, cols := 2, 3
	s := Constructor(rows, cols)

	recorder := make([]bool, rows*cols)
	for i := 0; i < rows*cols; i++ {
		f := s.Flip()
		r, c := f[0], f[1]
		old := recorder[r*cols+c]
		recorder[r*cols+c] = true
		ast.False(old)
	}

	s.Reset()

	recorder = make([]bool, rows*cols)
	for i := 0; i < rows*cols; i++ {
		f := s.Flip()
		r, c := f[0], f[1]
		old := recorder[r*cols+c]
		recorder[r*cols+c] = true
		ast.False(old)
	}

}
func Test_Solution_2(t *testing.T) {
	ast := assert.New(t)

	rows, cols := 1, 2
	s := Constructor(rows, cols)

	recorder := make([]bool, rows*cols)
	for i := 0; i < rows*cols; i++ {
		f := s.Flip()
		r, c := f[0], f[1]
		old := recorder[r*cols+c]
		recorder[r*cols+c] = true
		ast.False(old)
	}

	ast.Nil(s.Flip())
}
func Test_Solution_3(t *testing.T) {
	ast := assert.New(t)

	rows, cols := 10000, 10000
	s := Constructor(rows, cols)

	for times := 0; times < 200; times++ {
		recorder := make([]bool, rows*cols)
		s.Reset()
		for i := 0; i < 4; i++ {
			f := s.Flip()
			r, c := f[0], f[1]
			old := recorder[r*cols+c]
			recorder[r*cols+c] = true
			ast.False(old)
		}
	}

}
