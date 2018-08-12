package problem0705

// MyHashSet object will be instantiated and called as such:
// obj := Constructor();
// obj.Add(key);
// obj.Remove(key);
// param_3 := obj.Contains(key);
type MyHashSet struct {
	table []bool
}

// Constructor initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{table: make([]bool, 1000001)}
}

// Add 添加 key
func (s *MyHashSet) Add(key int) {
	s.table[key] = true
}

// Remove 移除 key
func (s *MyHashSet) Remove(key int) {
	s.table[key] = false
}

// Contains returns true if this set contains the specified element */
func (s *MyHashSet) Contains(key int) bool {
	return s.table[key]
}
