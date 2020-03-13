package problem0355

import "container/heap"

type Tweet struct {
	tweetId   int
	userId    int
	postion   int
	timeStamp int
}

type Twitter struct {
	tweets    map[int][]Tweet
	follows   map[int]map[int]bool
	timeStamp int
}

type MaxHeap []Tweet

func (this MaxHeap) Len() int {
	return len(this)
}

func (this MaxHeap) Less(i, j int) bool {
	return this[i].timeStamp > this[j].timeStamp
}

func (this MaxHeap) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this *MaxHeap) Push(x interface{}) {
	*this = append(*this, x.(Tweet))
}

func (this *MaxHeap) Pop() interface{} {
	tweet := (*this)[len(*this)-1]
	*this = (*this)[:len(*this)-1]

	return tweet
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{
		tweets:  make(map[int][]Tweet),
		follows: make(map[int]map[int]bool),
	}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.timeStamp++
	postion := len(this.tweets[userId])
	tweet := Tweet{tweetId, userId, postion, this.timeStamp}
	this.tweets[userId] = append(this.tweets[userId], tweet)
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	maxHeap := MaxHeap{}
	for followeeId := range this.follows[userId] {
		if len(this.tweets[followeeId]) == 0 {
			continue
		}

		maxHeap = append(maxHeap, this.tweets[followeeId][len(this.tweets[followeeId])-1])
	}

	if len(this.tweets[userId]) > 0 {
		maxHeap = append(maxHeap, this.tweets[userId][len(this.tweets[userId])-1])
	}

	heap.Init(&maxHeap)

	res := []int{}
	for i := 0; i < 10 && len(maxHeap) > 0; i++ {
		tweet := heap.Pop(&maxHeap).(Tweet)
		res = append(res, tweet.tweetId)
		if tweet.postion != 0 {
			heap.Push(&maxHeap, this.tweets[tweet.userId][tweet.postion-1])
		}
	}

	return res
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	}

	if len(this.follows[followerId]) == 0 {
		m := make(map[int]bool)
		this.follows[followerId] = m
	}

	if _, ok := this.follows[followerId][followeeId]; !ok {
		this.follows[followerId][followeeId] = true
	}
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	}

	delete(this.follows[followerId], followeeId)
}
