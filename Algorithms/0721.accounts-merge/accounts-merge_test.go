package problem0721

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	accounts [][]string
	ans      [][]string
}{

	{
		[][]string{
			{"David", "David0@m.co", "David1@m.co"},
			{"David", "David3@m.co", "David4@m.co"},
			{"David", "David4@m.co", "David5@m.co"},
			{"David", "David2@m.co", "David3@m.co"},
			{"David", "David1@m.co", "David2@m.co"},
		},
		[][]string{
			{"David", "David0@m.co", "David1@m.co", "David2@m.co", "David3@m.co", "David4@m.co", "David5@m.co"},
		},
	},

	{
		[][]string{
			{"Ethan", "Ethan1@m.co", "Ethan2@m.co", "Ethan0@m.co"},
			{"David", "David1@m.co", "David2@m.co", "David0@m.co"},
			{"Lily", "Lily0@m.co", "Lily0@m.co", "Lily4@m.co"},
			{"Gabe", "Gabe1@m.co", "Gabe4@m.co", "Gabe0@m.co"},
			{"Ethan", "Ethan2@m.co", "Ethan1@m.co", "Ethan0@m.co"},
		},
		[][]string{
			{"Gabe", "Gabe0@m.co", "Gabe1@m.co", "Gabe4@m.co"},
			{"Ethan", "Ethan0@m.co", "Ethan1@m.co", "Ethan2@m.co"},
			{"David", "David0@m.co", "David1@m.co", "David2@m.co"},
			{"Lily", "Lily0@m.co", "Lily4@m.co"},
		},
	},

	{
		[][]string{
			{"John", "johnsmith@mail.com", "john00@mail.com"},
			{"John", "johnnybravo@mail.com"},
			{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
			{"Mary", "mary@mail.com"},
		},
		[][]string{
			{"John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"},
			{"John", "johnnybravo@mail.com"},
			{"Mary", "mary@mail.com"},
		},
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(normal(tc.ans), normal(accountsMerge(tc.accounts)), "输入:%v", tc)
	}
}
func normal(sss [][]string) string {
	ss := make([]string, len(sss))

	for i := 0; i < len(sss); i++ {
		sort.Strings(sss[i])
		ss[i] = strings.Join(sss[i], "☆")
	}

	sort.Strings(ss)

	return strings.Join(ss, "★★")
}
func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			accountsMerge(tc.accounts)
		}
	}
}
