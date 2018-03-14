package problem0432

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AllOne(t *testing.T) {
	ast := assert.New(t)

	a := Constructor()

	ast.Equal("", a.GetMaxKey())
	ast.Equal("", a.GetMinKey())

	a.Inc("6")

	ast.Equal("6", a.GetMaxKey())
	ast.Equal("6", a.GetMinKey())

	a.Dec("6")

	ast.Equal("", a.GetMaxKey())
	ast.Equal("", a.GetMinKey())

	maxKeys := []string{"7", "8", "9"}
	minKeys := []string{"1", "2", "3"}

	for i := 0; i < 3; i++ {
		for j := 1; j < 10; j++ {
			a.Inc(strconv.Itoa(j))
		}
	}

	for _, key := range maxKeys {
		a.Inc(key)
		a.Inc(key)
	}

	for _, key := range minKeys {
		a.Dec(key)
		a.Dec(key)
	}

	ast.Contains(maxKeys, a.GetMaxKey(), "无法获取正确的 max key")

	ast.Contains(minKeys, a.GetMinKey(), "无法获取正确的 min key")

	// 以下是为了单元覆盖率
	a.Dec("100")
}
