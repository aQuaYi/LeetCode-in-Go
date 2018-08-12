package problem0705

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_myFunc(t *testing.T) {
	ast := assert.New(t)

	s := Constructor()

	s.Add(1)
	s.Add(2)

	ast.True(s.Contains(1))

	ast.False(s.Contains(3))

	s.Add(2)

	ast.True(s.Contains(2))

	s.Remove(2)

	ast.False(s.Contains(2))

}
