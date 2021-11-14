package question

// 实现一个 MapSum 类，支持两个方法，insert和sum：
// MapSum() 初始化 MapSum 对象
// void insert(String key, int val) 插入 key-val 键值对，字符串表示键 key ，整数表示值 val。
// 如果键 key 已经存在，那么原来的键值对将被替代成新的键值对。
// int sum(string prefix) 返回所有以该前缀 prefix 开头的键 key 的值的总和。

// 主要思想：要构建一棵前缀树，并记录子树的所有有效键值对的和
// 对插入的时候，因为可能涉及到修改已经存在键值对，所以要记录路径，如果是已经存在的键值对，需要调整路径上的sum
// 这里要严格注意的就是边界情况的处理！

type Node0677 struct {
	sum      int
	val      int
	children [26]*Node0677
}
type MapSum struct {
	root *Node0677
}

func Constructor() MapSum {
	return MapSum{
		root: &Node0677{
			sum:      0,
			val:      0,
			children: [26]*Node0677{},
		},
	}
}

func (this *MapSum) Insert(key string, val int) {
	p := this.root
	path := make([]*Node0677, len(key))
	cur := 0
	for cur < len(key) {
		path[cur] = p
		if p.children[key[cur]-'a'] == nil {
			p.children[key[cur]-'a'] = &Node0677{
				sum:      0,
				val:      0,
				children: [26]*Node0677{},
			}
		}
		p = p.children[key[cur]-'a']
		cur++
	}

	diff := 0
	if p.val == 0 {
		diff = val
	} else {
		// 说明是已经存在的键值对
		diff = val - p.val
	}
	p.val = val
	p.sum += diff
	for i := 0; i < len(key); i++ {
		path[i].sum += diff
	}
}

func (this *MapSum) Sum(prefix string) int {
	p := this.root
	cur := 0
	for cur < len(prefix) {
		if p.children[prefix[cur]-'a'] == nil {
			return 0
		}
		p = p.children[prefix[cur]-'a']
		cur++
	}
	return p.sum
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
