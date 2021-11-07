package question

// 给你一个长度为 n 的链表，每个节点包含一个额外增加的随机指针 random ，该指针可以指向链表中的任何节点或空节点。
//
//构造这个链表的深拷贝。深拷贝应该正好由 n 个 全新 节点组成，其中每个新节点的值都设为其对应的原节点的值。新节点的 next 指针和 random 指针也都应指向复制链表中的新节点，并使原链表和复制链表中的这些指针能够表示相同的链表状态。复制链表中的指针都不应指向原链表中的节点 。
//
//例如，如果原链表中有 X 和 Y 两个节点，其中 X.random --> Y 。那么在复制链表中对应的两个节点 x 和 y ，同样有 x.random --> y 。
//
//返回复制链表的头节点。
//
//用一个由n个节点组成的链表来表示输入/输出中的链表。每个节点用一个[val, random_index]表示：
//
//val：一个表示Node.val的整数。
//random_index：随机指针指向的节点索引（范围从0到n-1）；如果不指向任何节点，则为null。
//你的代码 只 接受原链表的头节点 head 作为传入参数。

type Node138 struct {
	Val    int
	Next   *Node138
	Random *Node138
}

func copyRandomList(head *Node138) *Node138 {
	return copyRandomListCore(head)
}

// 主要思想：为了避免随机无法找到，分为三步
// 首先借助原有的链表，复制所有的节点，这步无法复制随机指针
// 然后在复制所有节点的基础上，复制随机指针
// 最后将偶数位置上的节点拆下来
func copyRandomListCore(head *Node138) *Node138 {
	if head == nil {
		return nil
	}
	// 第一步，在节点之后复制一个新节点
	for node := head; node != nil; node = node.Next.Next {
		node.Next = &Node138{
			Val:    node.Val,
			Next:   node.Next,
			Random: nil,
		}
	}
	// 第二步，绑定新节点的随机指针
	for node := head; node != nil; node = node.Next.Next {
		if node.Random != nil {
			node.Next.Random = node.Random.Next
		}
	}
	// 第三步，拆出来
	ret := head.Next
	for node := head; node != nil; node = node.Next {
		t := node.Next
		node.Next = node.Next.Next
		if t.Next != nil {
			t.Next = t.Next.Next
		}
	}
	return ret
}
