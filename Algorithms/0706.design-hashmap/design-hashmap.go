package problem0706

// NOTICE: 120 ms 的答案，使用了内置的数据结构 map，违反了题目的要求

// MyHashMap object will be instantiated and called as such:
// obj := Constructor();
// obj.Put(key,value);
// param_2 := obj.Get(key);
// obj.Remove(key);
type MyHashMap struct {
	table []int
}

// Constructor initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{
		table: make([]int, 1000001),
	}
}

// Put value will always be non-negative. */
func (m *MyHashMap) Put(key int, value int) {
	// value + 1 是为了让 m.table 的默认值 0 和 value+1 的取值范围区分开。
	// m.table[key] = 0 就表示了 key 没有对应值
	// m.table[key] = 1 表示，key 有对应的值是 1-1 = 0
	m.table[key] = value + 1
}

// Get returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key
func (m *MyHashMap) Get(key int) int {
	return m.table[key] - 1
}

// Remove the mapping of the specified value key if this map contains a mapping for the key
func (m *MyHashMap) Remove(key int) {
	m.table[key] = 0
}
