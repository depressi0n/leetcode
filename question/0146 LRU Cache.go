package question

// 运用你所掌握的数据结构，设计和实现一个 LRU (最近最少使用) 缓存机制 。
// 实现 LRUCache 类：
// LRUCache(int capacity) 以正整数作为容量capacity 初始化 LRU 缓存
// int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
// void put(int key, int value)如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字-值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。
// 进阶：你是否可以在O(1) 时间复杂度内完成这两种操作？

// （1）双向链表存放实际存放的值，用map记录是否存储
// （2）可以将属于双向链表的操作独立出来，降低耦合度

type LRUCache struct {
	size       int                  //当前的size
	capacity   int                  //最大容量
	cache      map[int]*DLinkedNode //存放指针
	head, tail *DLinkedNode         //存放头尾
}

type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

func initDLinkedNode(key, value int) *DLinkedNode { //初始化一个炼表节点
	return &DLinkedNode{
		key:   key,
		value: value,
	}
}

func LRUConstructor(capacity int) LRUCache { //初始化函数
	l := LRUCache{
		cache:    map[int]*DLinkedNode{},
		head:     initDLinkedNode(0, 0), //虚节点
		tail:     initDLinkedNode(0, 0), //虚节点
		capacity: capacity,
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok { //首先是查询，查询失败就不管了
		return -1
	}
	this.moveToHead(node) //查询成功则将指针移到最前面
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.cache[key]
	if !ok { //首先也是查询，查询失败
		node = initDLinkedNode(key, value) //意味着要增加一个新节点
		this.cache[key] = node             //记录下来
		this.addToHead(node)               //将map放到head
		this.size++                        //size+1
		if this.size > this.capacity {     //如果超过容量，则去掉尾部节点
			removed := this.removeTail()
			delete(this.cache, removed.key)
			this.size-- //size--
		}
	} else { //如果查询成功
		node.value = value    //需要更新节点的value
		this.moveToHead(node) //并将节点移到最前面
	}
}

func (this *LRUCache) addToHead(node *DLinkedNode) { //添加一个元素到head
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DLinkedNode) { //移除给点节点
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *DLinkedNode) { //移动到head
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *DLinkedNode { //移除尾部到节点
	node := this.tail.prev
	this.removeNode(node)
	return node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
