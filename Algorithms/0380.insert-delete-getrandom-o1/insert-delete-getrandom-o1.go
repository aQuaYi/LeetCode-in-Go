package Problem0380

import "math/rand"

// RandomizedSet 是一个随机获取的集合
type RandomizedSet struct {
	rec map[int]bool
}

// Constructor returns RandomizedSet
/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		rec: make(map[int]bool),
	}
}

// Insert 插入数据
/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (r *RandomizedSet) Insert(val int) bool {
	if _, ok := r.rec[val]; ok {
		return false
	}
	r.rec[val] = true
	return true
}

// Remove 删除数据
/** Removes a value from the set. Returns true if the set contained the specified element. */
func (r *RandomizedSet) Remove(val int) bool {
	if _, ok := r.rec[val]; ok {
		delete(r.rec, val)
		return true
	}

	return false
}

// GetRandom 获取随机数据
/** Get a random element from the set. */
func (r *RandomizedSet) GetRandom() int {
	if len(r.rec) == 0 {
		return 0
	}

	rnd := rand.Int() % len(r.rec)
	res := 0
	for k := range r.rec {
		if rnd == 0 {
			res = k
			break
		}
		rnd--
	}

	return res
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
