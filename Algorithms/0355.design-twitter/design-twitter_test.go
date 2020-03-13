package problem0355

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Constructor(t *testing.T) {
	ast := assert.New(t)

	tw := Constructor()
	tw.PostTweet(1, 5)
	tw.GetNewsFeed(1)
	tw.Follow(1, 1)
	tw.Follow(1, 10)
	ast.Equal([]int{5}, tw.GetNewsFeed(1))

	tw.Follow(1, 2)
	tw.PostTweet(2, 6)
	tw.Follow(1, 2)
	ast.Equal([]int{6, 5}, tw.GetNewsFeed(1))

	tw.Unfollow(1, 2)
	ast.Equal([]int{5}, tw.GetNewsFeed(1))
	tw.Unfollow(1, 1)
	ast.Equal([]int{5}, tw.GetNewsFeed(1))

	tw.PostTweet(10, 50)
	tw.PostTweet(10, 51)
	ast.Equal([]int{51, 50, 5}, tw.GetNewsFeed(1))
}
