package problem0384

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Constructor(t *testing.T) {
	ast := assert.New(t)

	original := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}

	s := Constructor(original)

	ans := s.Shuffle()
	sort.Ints(ans)
	ast.Equal(original, ans)

	ans = s.Reset()
	ast.Equal(original, ans)

}
