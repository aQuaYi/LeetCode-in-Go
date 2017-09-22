package Problem0146

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
type node struct {
	prev, next *node
	key, val   int
}

// LRUCache is LRU 方法的具体实现
type LRUCache struct {
	len, cap    int
	first, last *node
	hash        []*node
}

// Constructor 创建容量为 capacity 的 cache
func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:  capacity,
		hash: make([]*node, 0, capacity),
	}
}

// Get 获取 cache 中的数据
func (c *LRUCache) Get(key int) int {

	return 0
}

// Put is 放入新数据
func (c *LRUCache) Put(key int, value int) {

}
