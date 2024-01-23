package top150

type LRUCache struct {
	capacity   int
	cache      map[int]*DLNode
	head, tail *DLNode
}

type DLNode struct {
	key, val  int
	pre, next *DLNode
}

func initDLNode(key int, value int) *DLNode {
	return &DLNode{key: key, val: value}
}

func Constructor(capacity int) LRUCache {
	c := LRUCache{
		capacity: capacity,
		cache:    map[int]*DLNode{},
		head:     initDLNode(0, 0),
		tail:     initDLNode(0, 0),
	}
	c.head.next = c.tail
	c.tail.pre = c.head
	return c
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.move2Head(node)
	return node.val
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.cache[key]
	if !ok {
		node = initDLNode(key, value)
		this.addHead(node)
		if len(this.cache) > this.capacity {
			nt := this.delTail()
			delete(this.cache, nt.key)
		}
		return
	}
	node.val = value
	this.move2Head(node)
}

func (this *LRUCache) move2Head(node *DLNode) {
	this.delNode(node)
	this.addHead(node)
}

func (this *LRUCache) addHead(node *DLNode) {
	node.next = this.head.next
	node.pre = this.head
	this.head.next.pre = node
	this.head.next = node
}

func (this *LRUCache) delNode(node *DLNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *LRUCache) delTail() *DLNode {
	node := this.tail.pre
	this.delNode(node)
	return node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
