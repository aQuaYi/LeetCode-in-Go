package problem0384

import (
	"sort"
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func Test_Constructor(t *testing.T) {
	ast := assert.New(t)

	original := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}

	s := Constructor(original)
	fmt.Println("-------------------------")
	fmt.Println(s)
	fmt.Println("-------------------------")
	fmt.Println("+++++++++++++++++++++++++")
	ans := s.Shuffle()
	fmt.Println(ans)
	fmt.Println("+++++++++++++++++++++++++")
	sort.Ints(ans)
	ast.Equal(original, ans)

	ans = s.Reset()
	ast.Equal(original, ans)

}
