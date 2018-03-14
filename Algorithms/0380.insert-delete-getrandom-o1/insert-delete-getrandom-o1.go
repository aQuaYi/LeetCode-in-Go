package problem0380

import "math/rand"

// RandomizedSet 是一个随机获取的集合
type RandomizedSet struct {
	a   []int
	idx map[int]int // idx 的 v 是 k 在 a 中的位置。
}

// Constructor returns RandomizedSet
/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		a:   []int{},
		idx: make(map[int]int),
	}
}

// Insert 插入数据
/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (r *RandomizedSet) Insert(val int) bool {
	if _, ok := r.idx[val]; ok {
		return false
	}
	r.a = append(r.a, val)
	r.idx[val] = len(r.a) - 1
	return true
}

// Remove 删除数据
/** Removes a value from the set. Returns true if the set contained the specified element. */
func (r *RandomizedSet) Remove(val int) bool {
	if _, ok := r.idx[val]; !ok {
		return false
	}

	// 把 a 的最后一个数，放入待删除的数的位置
	r.a[r.idx[val]] = r.a[len(r.a)-1]
	r.idx[r.a[len(r.a)-1]] = r.idx[val]
	// 删除最后一个数
	r.a = r.a[:len(r.a)-1]
	// 在 r.idx 中删除 val 的记录
	delete(r.idx, val)

	return true
}

// GetRandom 获取随机数据
/** Get a random element from the set. */
func (r *RandomizedSet) GetRandom() int {
	return r.a[rand.Intn(len(r.a))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
