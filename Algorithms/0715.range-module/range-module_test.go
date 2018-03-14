package problem0715

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	rm := Constructor()

	rm.AddRange(10, 20)

	rm.RemoveRange(14, 16)

	ast.True(rm.QueryRange(10, 14))

	ast.False(rm.QueryRange(13, 15))

	ast.True(rm.QueryRange(16, 17))
}

func Test_RangeModule_AddRange(t *testing.T) {
	ast := assert.New(t)

	rm := Constructor()

	ranges := [][]int{
		{10, 12},
		{14, 16},
		{18, 20},
		{12, 14},
		{16, 18},
	}

	for _, r := range ranges {
		rm.AddRange(r[0], r[1])
	}

	ast.Equal(1, len(rm.ranges), "添加完 ranges 后，rm.ranges 的范围应该合并成一个")
	ast.Equal(interval{left: 10, right: 20}, *rm.ranges[0], "添加完 ranges 后，rm.ranges 的范围应该合并成[10,20)")
}

func Test_RangeModule_RemoveRange(t *testing.T) {
	ast := assert.New(t)

	rm := Constructor()

	rm.RemoveRange(1, 2)

	ranges := [][]int{
		{10, 12},
		{14, 16},
	}

	for _, r := range ranges {
		rm.AddRange(r[0], r[1])
	}

	rm.RemoveRange(12, 14)

	ast.Equal(2, len(rm.ranges), "添加完 ranges 后，rm.ranges 的范围应该合并成一个")
	ast.Equal(interval{left: 10, right: 12}, *rm.ranges[0])
	ast.Equal(interval{left: 14, right: 16}, *rm.ranges[1])
}

func Test_minus(t *testing.T) {
	ast := assert.New(t)

	var tcs = []struct {
		a, b       *interval
		ansA, ansB *interval
	}{

		{
			&interval{left: 10, right: 20},
			&interval{left: 5, right: 25},
			nil,
			nil,
		},

		{
			&interval{left: 10, right: 20},
			&interval{left: 5, right: 15},
			&interval{left: 15, right: 20},
			nil,
		},

		{
			&interval{left: 10, right: 20},
			&interval{left: 15, right: 25},
			&interval{left: 10, right: 15},
			nil,
		},

		{
			&interval{left: 10, right: 20},
			&interval{left: 14, right: 16},
			&interval{left: 10, right: 14},
			&interval{left: 16, right: 20},
		},
	}

	for _, tc := range tcs {
		ra, rb := minus(tc.a, tc.b)
		ast.Equal(tc.ansA, ra)
		ast.Equal(tc.ansB, rb)
	}
}
