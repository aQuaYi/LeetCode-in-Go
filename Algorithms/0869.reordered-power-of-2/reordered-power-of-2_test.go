package problem0869

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	N   int
	ans bool
}{

	{
		160,
		false,
	},

	{
		1,
		true,
	},

	{
		10,
		false,
	},

	{
		16,
		true,
	},

	{
		24,
		false,
	},

	{
		46,
		true,
	},

	// 可以有多个 testcase
}

func Test_reorderedPowerOf2(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, reorderedPowerOf2(tc.N), "输入:%v", tc)
	}
}

func Benchmark_reorderedPowerOf2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			reorderedPowerOf2(tc.N)
		}
	}
}

var val = 9876543210

func Benchmark_encode(b *testing.B) {
	for i := 1; i < b.N; i++ {
		encode(val)
	}
}

func encodeString(n int) string {
	tmp := [10]int{}
	for n > 0 {
		tmp[n%10]++
		n /= 10
	}
	ss := make([]string, 0, 10)
	for n, c := range tmp {
		if c > 0 {
			ss = append(ss, fmt.Sprintf("%d_%d", n, c))
		}
	}
	return strings.Join(ss, "-")
}

func Benchmark_encodeString(b *testing.B) {
	for i := 1; i < b.N; i++ {
		encodeString(val)
	}
}
