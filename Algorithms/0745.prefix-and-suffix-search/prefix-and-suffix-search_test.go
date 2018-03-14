package problem0745

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_wordFilter(t *testing.T) {
	ast := assert.New(t)

	words := []string{"apple", "banana", "abbbe"}
	w := Constructor(words)

	ast.Equal(0, w.F("ap", "e"))

	ast.Equal(1, w.F("ba", ""))

	ast.Equal(2, w.F("a", "e"))

	ast.Equal(2, w.F("", ""))
}
func Test_wordFilter_2(t *testing.T) {
	ast := assert.New(t)

	words := []string{"pop"}
	w := Constructor(words)

	pss := [][]string{
		{"", ""}, {"", "p"}, {"", "op"}, {"", "pop"}, {"p", ""}, {"p", "p"}, {"p", "op"}, {"p", "pop"}, {"po", ""}, {"po", "p"}, {"po", "op"}, {"po", "pop"}, {"pop", ""}, {"pop", "p"}, {"pop", "op"}, {"pop", "pop"}, {"", ""}, {"", "p"}, {"", "gp"}, {"", "pgp"}, {"p", ""}, {"p", "p"}, {"p", "gp"}, {"p", "pgp"}, {"pg", ""}, {"pg", "p"}, {"pg", "gp"}, {"pg", "pgp"}, {"pgp", ""}, {"pgp", "p"}, {"pgp", "gp"}, {"pgp", "pgp"}}

	ans := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, -1, -1, 0, 0, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}

	for i := 0; i < len(pss); i++ {
		ast.Equal(ans[i], w.F(pss[i][0], pss[i][1]), "输入 %v, %d", pss[i], ans[i])
	}
}
