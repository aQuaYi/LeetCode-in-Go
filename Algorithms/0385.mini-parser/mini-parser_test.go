package problem0385

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s   string
	ans *NestedInteger
}{

	{"", nil},

	{"324", &NestedInteger{Num: 324}},

	{"[123,[456,[789]]]",
		&NestedInteger{
			Ns: []*NestedInteger{
				&NestedInteger{Num: 123},
				&NestedInteger{
					Ns: []*NestedInteger{
						&NestedInteger{Num: 456},
						&NestedInteger{
							Ns: []*NestedInteger{
								&NestedInteger{Num: 789},
							},
						},
					},
				},
			},
		},
	},

	// 可以有多个 testcase
}

func Test_deserialize(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, deserialize(tc.s), "输入:%v", tc)
	}
}
