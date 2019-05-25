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
	if _, ok := m.times[key]; !ok {
		m.times[key] = make([]int, 1, 1024)
		m.strings[key] = make([]string, 1, 1024)
	}
	m.times[key] = append(m.times[key], timestamp)
	m.strings[key] = append(m.strings[key], value)
}

// Get the value of key
func (m *TimeMap) Get(key string, timestamp int) string {
	times := m.times[key]
	i := sort.Search(len(times), func(i int) bool {
		return timestamp < times[i]
	})
	i--
	return m.strings[key][i]
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
