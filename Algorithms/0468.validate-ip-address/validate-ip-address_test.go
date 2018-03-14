package problem0468

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	IP  string
	ans string
}{

	{
		"1081:db8:85a3:01:-0:8A2E:0370:7334",
		"Neither",
	},

	{
		"2001:0db8:85a3:0000:0:8A2E:0370:733a",
		"IPv6",
	},

	{
		"0.0.0.-0",
		"Neither",
	},

	{
		"2001:0db8:85a3:0:0:8A2E:0370:7334",
		"IPv6",
	},

	{
		"172.16.254.1",
		"IPv4",
	},

	{
		"0db8:85a3:00000:0:8A2E:0370:7334",
		"Neither",
	},

	{
		"2001:0db8:85a3:00000:0:8A2E:0370:7334",
		"Neither",
	},

	{
		"56.256.256",
		"Neither",
	},

	{
		"25.25..25",
		"Neither",
	},

	{
		"256.256.256.256",
		"Neither",
	},

	// 可以有多个 testcase
}

func Test_validIPAddress(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, validIPAddress(tc.IP), "输入:%v", tc)
	}
}

func Benchmark_validIPAddress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			validIPAddress(tc.IP)
		}
	}
}
