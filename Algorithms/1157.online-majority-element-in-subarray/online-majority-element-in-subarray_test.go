package problem1157

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MajorityChecker(t *testing.T) {
	a := assert.New(t)
	//
	arr := []int{1, 1, 2, 2, 1, 1}
	mc := Constructor(arr)
	tcs := []struct {
		left, right, threshold, ans int
	}{
		{0, 5, 4, 1},
		{0, 3, 3, -1},
		{2, 3, 2, 2},
	}
	for _, tc := range tcs {
		a.Equal(tc.ans, mc.Query(tc.left, tc.right, tc.threshold))
	}
}
