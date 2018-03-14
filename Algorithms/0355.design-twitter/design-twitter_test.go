package problem0355

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Constructor(t *testing.T) {
	ast := assert.New(t)
	twitter := Constructor()

	twitter.PostTweet(1, 5)
	twitter.Follow(1, 1)

	ast.Equal([]int{5}, twitter.GetNewsFeed(1))

	twitter.Follow(1, 2)

	twitter.PostTweet(2, 4)

	ast.Equal([]int{4, 5}, twitter.GetNewsFeed(1))

	twitter.Follow(1, 2)
	twitter.Unfollow(1, 2)

	ast.Equal([]int{5}, twitter.GetNewsFeed(1))
}
