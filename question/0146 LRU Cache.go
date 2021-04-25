package question

//双向链表存放实际存放的值，用map记录是否存储
// 可以将属于双向链表的操作独立出来，降低耦合度
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
