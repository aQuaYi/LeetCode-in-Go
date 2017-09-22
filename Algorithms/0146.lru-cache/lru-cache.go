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
	nodes       map[int]*node
}

// Constructor 创建容量为 capacity 的 cache
func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap: capacity,
		// TODO: 检查这样设置的必要性
		nodes: make(map[int]*node, capacity),
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

func (c *LRUCache) remove(key int) {
	nd, ok := c.nodes[key]
	if ok {
		if nd.prev != nil {
			nd.prev.next = nd.next
		}

		if nd.next != nil {
			nd.next.prev = nd.prev
		}

		if c.last == nd {
			c.last = nd.prev
		}

		if nd == c.first {
			c.first = nd.next
		}
	}
}

func (c *LRUCache) removeLast() {
	if c.last == nil {
		return
	}

	if c.last.prev != nil {
		c.last.prev.next = nil
	} else {
		c.first = nil
	}

	c.last = c.last.prev
}

func (c *LRUCache) moveToHead(nd *node) {
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
		if c.len >= c.cap {
			// TODO: 是否可以删除这个 if
			if c.last != nil {
				delete(c.nodes, c.last.key)
				c.removeLast()
			}
		} else {
			c.len++
		}
		nd = &node{}
	}

	nd.val = value
	nd.key = key
	c.moveToHead(nd)
	c.nodes[key] = nd
}
