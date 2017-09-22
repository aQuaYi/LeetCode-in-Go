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
	// 利用 hashTable 查询 key
	if node, ok := c.nodes[key]; ok {
		// key 存在的话
		// 把对应的 node 移动到 cache 的双向链的 first 位
		c.moveToFirst(node)
		return node.val
	}

	// key 不存在，按照题意，返回 -1
	return -1
}

// Put is 放入新数据
func (c *LRUCache) Put(key int, value int) {
	node, ok := c.nodes[key]

	if ok { // 更新旧 node
		// 更新 node 中的值
		node.val = value
		// 把 node 放入双向链的 first 位
		c.moveToFirst(node)
	} else { // 放入新 node
		if c.len == c.cap {
			// cache 已满，删除 last 位，为新 node 腾地方
			// 删除 hashTable 中的记录
			delete(c.nodes, c.last.key)
			// 删除双向链中的 last 位
			c.removeLast()
		} else {
			c.len++
		}
		// 为 key 和 value 新建一个节点
		node = &doublyLinkedNode{
			key: key,
			val: value,
		}
		// 在 hashTable 中添加记录
		c.nodes[key] = node
		// 把新 node 放入 first 位
		c.insertToFirst(node)
	}
}

func (c *LRUCache) removeLast() {
	if c.last.prev != nil { // 双向链长度 >1
		c.last.prev.next = nil
	} else { // 双向链长度 == 1，firt，last 指向同一个节点
		c.first = nil
	}

	c.last = c.last.prev
}

func (c *LRUCache) moveToFirst(node *doublyLinkedNode) {
	switch node {
	case c.first:
		return
	case c.last:
		c.removeLast()
	default:
		// 在双向链中，删除 node 节点
		node.prev.next = node.next
		node.next.prev = node.prev
	}

	// 策略是
	// 如果需要移动 node 的话
	// 先删除，再插入
	c.insertToFirst(node)
}

func (c *LRUCache) insertToFirst(node *doublyLinkedNode) {
	if c.last == nil { // **空**双向链
		c.last = node
	} else { // **非空**双向链
		// 默认 node != nil
		node.next = c.first
		c.first.prev = node
	}

	c.first = node
}
