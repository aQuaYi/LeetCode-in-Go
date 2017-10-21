package Problem0355

// Twitter is twitter user
type Twitter struct {
}

// Constructor initialize your data structure here.
func Constructor() Twitter {

	return Twitter{}
}

// PostTweet compose a new tweet.
func (t *Twitter) PostTweet(userID int, tweetID int) {

}

// GetNewsFeed retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent.
func (t *Twitter) GetNewsFeed(userID int) []int {

	return []int{}
}

// Follow followee. If the operation is invalid, it should be a no-op.
func (t *Twitter) Follow(followerID int, followeeID int) {

}

// Unfollow follower unfollows a followee. If the operation is invalid, it should be a no-op.
func (t *Twitter) Unfollow(followerID int, followeeID int) {

}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userID,tweetID);
 * param_2 := obj.GetNewsFeed(userID);
 * obj.Follow(followerID,followeeID);
 * obj.Unfollow(followerID,followeeID);
 */
