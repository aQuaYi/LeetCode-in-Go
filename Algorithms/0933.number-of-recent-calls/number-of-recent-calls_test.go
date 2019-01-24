package problem0933

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_myFunc(t *testing.T) {
	ast := assert.New(t)

	rc := Constructor()

	times := []int{1, 100, 3001, 3002}
	counters := []int{1, 2, 3, 3}

	for i := range times {
		t, c := times[i], counters[i]
		ast.Equal(c, rc.Ping(t))
	}
}
