package problem0677

import "strings"
import "sort"

// MapSum 可以返回　map 中具有相同前缀的　value 之和
type MapSum struct {
	m    map[string]int
	keys []string
}

// Constructor 构建了 MapSum
func Constructor() MapSum {
	return MapSum{m: make(map[string]int, 128),
		keys: make([]string, 0, 128)}
}

// Insert 插入 (key, val)
func (ms *MapSum) Insert(key string, val int) {
	ms.m[key] = val

	i := sort.SearchStrings(ms.keys, key)
	// 遇到重复的 key ，不操作
	if i < len(ms.keys) && ms.keys[i] == key {
		return
	}

	ms.keys = append(ms.keys, "")
	copy(ms.keys[i+1:], ms.keys[i:])
	ms.keys[i] = key
}

// Sum 返回所有具有 prefix 的 key 的 val 之和
func (ms *MapSum) Sum(prefix string) int {
	res := 0
	i := sort.SearchStrings(ms.keys, prefix)
	for i < len(ms.keys) && strings.HasPrefix(ms.keys[i], prefix) {
		res += ms.m[ms.keys[i]]
		i++
	}
	return res
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
