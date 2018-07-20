package problem0470

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rand10(t *testing.T) {
	ast := assert.New(t)
	for i := 0; i < 100; i++ {
		r := rand10()
		ast.True(1 <= r && r <= 10)
	}
}

func Benchmark_rand10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand10()
	}
}
