package problem0981

import "sort"

// TimeMap is a key-value store with timestamp
type TimeMap struct {
	times   map[string][]int
	strings map[string][]string
}

// Constructor Initialize your data structure here.
func Constructor() TimeMap {
	return TimeMap{
		times:   make(map[string][]int, 1024),
		strings: make(map[string][]string, 1024),
	}
}

// Set the key and the value
// Note 3:
// The timestamps for all TimeMap.set operations are strictly increasing.
func (m *TimeMap) Set(key string, value string, timestamp int) {
	m.times[key] = append(m.times[key], timestamp)
	m.strings[key] = append(m.strings[key], value)
}

// Get the value of key
func (m *TimeMap) Get(key string, timestamp int) string {
	if len(m.times[key]) == 0 ||
		timestamp < m.times[key][0] {
		return ""
	}
	i := sort.SearchInts(m.times[key], timestamp)
	if i == len(m.times[key]) ||
		timestamp < m.times[key][i] {
		i--
	}
	return m.strings[key][i]
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
