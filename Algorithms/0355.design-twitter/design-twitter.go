package problem0355

import "sort"
import "time"

type tweet struct {
	id   int
	time int64
}

// tweets 用于排序
type tweets []tweet

func (t tweets) Len() int {
	return len(t)
}
func (t tweets) Less(i, j int) bool {
	return t[i].time > t[j].time
}
func (t tweets) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

// Twitter is twitter user
type Twitter struct {
	userTweets map[int]tweets
	follow     map[int][]int
}

// Constructor initialize your data structure here.
func Constructor() Twitter {
	t := make(map[int]tweets)
	f := make(map[int][]int)
	return Twitter{userTweets: t, follow: f}
}

// PostTweet compose a new tweet.
func (t *Twitter) PostTweet(userID int, tweetID int) {
	t.userTweets[userID] = append(
		t.userTweets[userID],
		tweet{
			id:   tweetID,
			time: time.Now().UnixNano(),
		},
	)
}

// GetNewsFeed retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent.
func (t *Twitter) GetNewsFeed(userID int) []int {
	// 获取本人的 tweets
	temp := make(tweets, len(t.userTweets[userID]))
	copy(temp, t.userTweets[userID])
	// 获取 followee 的 tweets
	for _, id := range t.follow[userID] {
		temp = append(temp, t.userTweets[id]...)
	}
	// 按照时间排序
	sort.Sort(temp)
	// 获取最近的 10 条或更少的内容
	res := make([]int, 0, 10)
	for i := 0; i < len(temp) && i < 10; i++ {
		res = append(res, temp[i].id)
	}
	return res
}

// Follow followee. If the operation is invalid, it should be a no-op.
func (t *Twitter) Follow(followerID int, followeeID int) {
	// 不能 follow 自己
	if followerID == followeeID {
		return
	}
	// 不能重复 follow
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
			// 删除 followeeID 记录
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
