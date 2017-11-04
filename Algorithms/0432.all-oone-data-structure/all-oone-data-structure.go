package Problem0432

// AllOne 是解题所需的结构
type AllOne struct {
	m              map[string]int
	max, min       int
	maxKey, minKey string
}

// Constructor initialize your data structure here.
func Constructor() AllOne {
	return AllOne{m: make(map[string]int),
		max: -1 << 63,
		min: 1<<63 - 1,
	}
}

// Inc inserts a new key <Key> with value 1. Or increments an existing key by 1.
func (a *AllOne) Inc(key string) {
	if _, ok := a.m[key]; ok {
		a.m[key]++
		if a.max < a.m[key] {
			a.max = a.m[key]
			a.maxKey = key
		}
	} else {
		a.m[key] = 1
		if a.min > 1 {
			a.min = 1
			a.minKey = key
		}
	}

}

// Dec decrements an existing key by 1. If Key's value is 1, remove it from the data structure.
func (a *AllOne) Dec(key string) {
	if _, ok := a.m[key]; !ok {
		return
	}

	a.m[key]--

	if a.min > a.m[key] {
		a.min = a.m[key]
		a.minKey = key
	}
}

// GetMaxKey returns one of the keys with maximal value.
func (a *AllOne) GetMaxKey() string {
	return a.maxKey
}

// GetMinKey returns one of the keys with Minimal value.
func (a *AllOne) GetMinKey() string {
	return a.minKey
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
