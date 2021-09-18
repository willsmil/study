package arithmetic

type LRCCache struct {
	m    map[string]*Node
	cap  int
	head *Node
	tail *Node
}

type Node struct {
	pre   *Node
	next  *Node
	value interface{}
}

func (c *LRCCache) Get(key string) interface{} {
	if v, ok := c.m[key]; ok {
		c.MoveToHead(v)
		return v.value
	}
	return nil
}

func (c *LRCCache) Set(key string, value interface{}) {
	if v, ok := c.m[key]; ok {
		c.m[key].value = value
		// 移动到最前面
		c.MoveToHead(v)
		return
	}
	node := Node{pre: nil, next: nil, value: value}
	if len(c.m) < c.cap {
		c.m[key] = &node
		c.AddHead(&node)
		return
	}
	// 放最前面
	c.m[key] = &node
	c.RemoveLast()
	c.AddHead(&node)
}

func (c *LRCCache) MoveToHead(n *Node) {
	c.Remove(n)
	c.AddHead(n)
}

func (c *LRCCache) Remove(n *Node) {
	pre := n.pre
	next := n.next
	pre.next = next
	next.pre = pre
}

func (c *LRCCache) RemoveLast() {
	c.Remove(c.tail.pre)
}

func (c *LRCCache) AddHead(n *Node) {
	n.next = c.head.next
	n.next.pre = n
	n.pre = c.head
	c.head.next = n
}
