package problem0911

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_myFunc(t *testing.T) {
	ast := assert.New(t)

	persons := []int{0, 1, 1, 0, 0, 1, 0}
	times := []int{0, 5, 10, 15, 20, 25, 30}

	qs := []int{3, 12, 25, 15, 24, 8}
	as := []int{0, 1, 1, 0, 0, 1}

	tvc := Constructor(persons, times)

	for i, q := range qs {
		expected := tvc.Q(q)
		actual := as[i]

		ast.Equal(expected, actual)
	}
}
