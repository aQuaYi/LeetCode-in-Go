package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_url(t *testing.T) {
	expected := "https://leetcode.com/api/problems/algorithms/"
	key := "algorithms"
	assert.Equal(t, expected, url(key), "生成的url不对")
}
