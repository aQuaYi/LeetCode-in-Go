package Problem0432

import "math/rand"

type item struct {
	key   string
	value int
	index int
}

// AllOne 是解题所需的结构
type AllOne struct {
	m                  map[string]*item
	max, min           int
	maxItems, minItems []*item
}

// Constructor initialize your data structure here.
func Constructor() AllOne {
	return AllOne{m: make(map[string]*item),
		max: -1 << 63,
		min: 1<<63 - 1,
	}
}

// Inc inserts a new key <Key> with value 1. Or increments an existing key by 1.
func (a *AllOne) Inc(key string) {
	if _, ok := a.m[key]; ok {
		a.m[key].value++
	} else {
		a.m[key] = &item{key: key, value: 1}
	}

	if a.m[key].value < a.min {
		a.min = a.m[key].value
		a.minItems = []*item{a.m[key]}
		a.m[key].index = 0
	} else if a.m[key].value == a.min {
		a.minItems = append(a.minItems, a.m[key])
		a.m[key].index = len(a.minItems) - 1
	} else if a.max == a.m[key].value {
		a.maxItems = append(a.maxItems, a.m[key])
		a.m[key].index = len(a.maxItems) - 1
	} else if a.max < a.m[key].value {
		a.max = a.m[key].value
		a.maxItems = []*item{a.m[key]}
		a.m[key].index = 0
	}
}

// Dec decrements an existing key by 1. If Key's value is 1, remove it from the data structure.
func (a *AllOne) Dec(key string) {
	if _, ok := a.m[key]; !ok {
		return
	}

	a.m[key].value--

	if a.m[key].value < a.min {
		a.min = a.m[key].value
		a.minItems = []*item{a.m[key]}
		a.m[key].index = 0
	} else if a.min == a.m[key].value {
		a.minItems = append(a.minItems, a.m[key])
		a.m[key].index = len(a.minItems) - 1
	}
}

// GetMaxKey returns one of the keys with maximal value.
func (a *AllOne) GetMaxKey() string {
	if len(a.maxItems) == 0 {
		return ""
	}
	return a.maxItems[rand.Intn(len(a.maxItems))].key
}

// GetMinKey returns one of the keys with Minimal value.
func (a *AllOne) GetMinKey() string {
	if len(a.minItems) == 0 {
		return ""
	}
	return a.minItems[rand.Intn(len(a.minItems))].key
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
