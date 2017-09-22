package Problem0146

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

// doublyLinkedNode 是双向链节点
type doublyLinkedNode struct {
	prev, next *doublyLinkedNode
	key, val   int
}

// LRUCache 利用 双向链条 + hashtabl 实现
type LRUCache struct {
	// cache 的 长度 和 容量
	len, cap int
	// 分别指向双向链的首尾节点
	first, last *doublyLinkedNode
	// 节点的 hashTable，方便查找节点是否存在
	nodes map[int]*doublyLinkedNode
}

// Constructor 创建容量为 capacity 的 cache
func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:   capacity,
		nodes: make(map[int]*doublyLinkedNode, capacity),
	}
}

// Get 获取 cache 中的数据
func (c *LRUCache) Get(key int) int {
	if node, ok := c.nodes[key]; ok {
		c.moveToHead(node)
		return node.val
	}
	return -1
}

func (c *LRUCache) removeLast() {

	if c.last.prev != nil {
		c.last.prev.next = nil
	} else {
		c.first = nil
	}

	c.last = c.last.prev
}

func (c *LRUCache) moveToHead(nd *doublyLinkedNode) {
	if nd == c.first {
		return
	}

	if nd.prev != nil {
		nd.prev.next = nd.next
	}

	if nd.next != nil {
		nd.next.prev = nd.prev
	}

	if nd == c.last {
		c.last = nd.prev
	}
	if c.first != nil {
		nd.next = c.first
		c.first.prev = nd
	}

	c.first = nd
	nd.prev = nil
	if c.last == nil {
		c.last = c.first
	}
}

// Put is 放入新数据
func (c *LRUCache) Put(key int, value int) {
	nd, ok := c.nodes[key]

	if !ok {
		if c.len == c.cap {
			delete(c.nodes, c.last.key)
			c.removeLast()
		} else {
			c.len++
		}
		nd = &doublyLinkedNode{}
	}

	nd.val = value
	nd.key = key
	c.moveToHead(nd)
	c.nodes[key] = nd
}
