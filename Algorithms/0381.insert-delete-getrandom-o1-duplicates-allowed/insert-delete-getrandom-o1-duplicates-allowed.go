package Problem0381

import "math/rand"

// RandomizedCollection 是一个随机获取的集合
type RandomizedCollection struct {
	a   []int
	idx map[int][]int // idx 的 v 是 k 在 a 中的位置。
}

// Constructor returns RandomizedSet
/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	return RandomizedCollection{
		a:   []int{},
		idx: make(map[int][]int),
	}
}

// Insert 插入数据
/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (r *RandomizedCollection) Insert(val int) bool {
	res := false
	if _, ok := r.idx[val]; !ok {
		res = true
	}
	r.a = append(r.a, val)
	r.idx[val] = append(r.idx[val], len(r.a)-1)
	return res
}

// Remove 删除数据
/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (r *RandomizedCollection) Remove(val int) bool {
	if _, ok := r.idx[val]; !ok {
		return false
	}

	last := r.a[len(r.a)-1]
	lenOfLast := len(r.idx[last])
	lenOfVal := len(r.idx[val])
	// 把 a 的最后一个数，放入最后一个 val 的位置
	r.a[r.idx[val][lenOfVal-1]] = last
	r.idx[last][lenOfLast-1] = r.idx[val][lenOfVal-1]

	// 删除最后一个数
	r.a = r.a[:len(r.a)-1]
	// 在 r.idx 中删除 val 的记录
	if lenOfVal == 1 {
		delete(r.idx, val)
	} else {
		r.idx[val] = r.idx[val][:lenOfVal-1]
	}

	return true
}

// GetRandom 获取随机数据
/** Get a random element from the collection. */
func (r *RandomizedCollection) GetRandom() int {
	return r.a[rand.Intn(len(r.a))]
}
