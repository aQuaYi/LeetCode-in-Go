package Problem0355

import (
	"sort"
)

// Twitter is twitter user
type Twitter struct {
	tweets map[int][]int
	follow map[int][]int
}

// Constructor initialize your data structure here.
func Constructor() Twitter {
	t := make(map[int][]int)
	f := make(map[int][]int)
	return Twitter{tweets: t, follow: f}
}

// PostTweet compose a new tweet.
func (t *Twitter) PostTweet(userID int, tweetID int) {
	t.tweets[userID] = append(t.tweets[userID], tweetID)
}

// GetNewsFeed retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent.
func (t *Twitter) GetNewsFeed(userID int) []int {
	res := make([]int, len(t.tweets[userID]))
	copy(res, t.tweets[userID])
	for _, id := range t.follow[userID] {
		res = append(res, t.tweets[id]...)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(res)))
	if len(res) > 10 {
		return res[:10]
	}
	return res
}

// Follow followee. If the operation is invalid, it should be a no-op.
func (t *Twitter) Follow(followerID int, followeeID int) {
	for _, id := range t.follow[followerID] {
		if id == followeeID {
			return
		}
	}
	t.follow[followerID] = append(t.follow[followerID], followeeID)
}

// Unfollow follower unfollows a followee. If the operation is invalid, it should be a no-op.
func (t *Twitter) Unfollow(followerID int, followeeID int) {
	for i, id := range t.follow[followerID] {
		if id == followeeID {
			t.follow[followerID] = append(t.follow[followerID][:i], t.follow[followerID][i+1:]...)
		}
	}
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userID,tweetID);
 * param_2 := obj.GetNewsFeed(userID);
 * obj.Follow(followerID,followeeID);
 * obj.Unfollow(followerID,followeeID);
 */
