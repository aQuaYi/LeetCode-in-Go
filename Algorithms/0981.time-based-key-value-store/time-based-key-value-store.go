package problem0981

import "sort"

type data struct {
	time  int
	value string
}

// TimeMap is a timebased key-value store
type TimeMap map[string][]data

// Constructor Initialize your data structure here.
func Constructor() TimeMap {
	return make(map[string][]data, 1024)
}

// Set the key and the value
// Note 3:
// The timestamps for all TimeMap.set operations are strictly increasing.
func (m TimeMap) Set(key string, value string, timestamp int) {
	if _, ok := m[key]; !ok {
		m[key] = make([]data, 1, 1024)
	}
	m[key] = append(m[key], data{
		time:  timestamp,
		value: value,
	})
}

// Get the value of key
func (m TimeMap) Get(key string, timestamp int) string {
	d := m[key]
	i := sort.Search(len(d), func(i int) bool {
		return timestamp < d[i].time
	})
	i--
	return m[key][i].value
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
