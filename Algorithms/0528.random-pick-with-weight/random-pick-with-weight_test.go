package problem0528

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Solution(t *testing.T) {
	ast := assert.New(t)

	w := []int{1, 3}
	total := 0.
	for i := range w {
		total += float64(w[i])
	}

	s := Constructor(w)
	count := make([]int, len(w))

	picks := 10000
	for i := 0; i < picks; i++ {
		index := s.PickIndex()
		count[index]++
	}

	for i, c := range count {
		expected := float64(w[i]) / total
		actual := float64(c) / float64(picks)
		delta := 0.001
		ast.InDelta(expected, actual, delta)
	}

}
func Test_Solution_2(t *testing.T) {
	ast := assert.New(t)

	w := []int{1, 3, 5, 7, 9, 11, 13, 15}
	total := 0.
	for i := range w {
		total += float64(w[i])
	}

	s := Constructor(w)
	count := make([]int, len(w))

	picks := 1000000
	for i := 0; i < picks; i++ {
		index := s.PickIndex()
		count[index]++
	}

	for i, c := range count {
		expected := float64(w[i]) / total
		actual := float64(c) / float64(picks)
		delta := 0.001
		ast.InDelta(expected, actual, delta)
	}

}
