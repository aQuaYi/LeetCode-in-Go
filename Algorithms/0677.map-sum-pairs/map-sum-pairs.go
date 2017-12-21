package Problem0677

import (
	"strings"
)

// MapSum 可以返回　map 中具有相同前缀的　value 之和
type MapSum struct {
	m map[string]int
}

// Constructor 构建了 MapSum
func Constructor() MapSum {
	return MapSum{m: make(map[string]int, 128)}
}

// Insert 插入 (key, val)
func (ms *MapSum) Insert(key string, val int) {
	ms.m[key] = val
}

// Sum 返回所有具有 prefix 的 key 的 val 之和
func (ms *MapSum) Sum(prefix string) int {
	res := 0
	for k, v := range ms.m {
		if strings.HasPrefix(k, prefix) {
			res += v
		}
	}
	return res
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
