package question

import (
	"math"
)

// 对链表进行插入排序。
//插入排序的动画演示如上。从第一个元素开始，该链表可以被认为已经部分排序（用黑色表示）。
//每次迭代时，从输入数据中移除一个元素（用红色表示），并原地将其插入到已排好序的链表中。
//插入排序算法：
//插入排序是迭代的，每次只移动一个元素，直到所有元素可以形成一个有序的输出列表。
//每次迭代中，插入排序只从输入数据中移除一个待排序的元素，找到它在序列中适当的位置，并将其插入。
//重复直到所有输入数据插入完为止。
func insertionSortList(head *ListNode) *ListNode {
	return insertionSortListCore(head)
}

// 主要思想：按照插入排序的思路，借助一个虚拟的头节点作为已经排序好的结果
// 到来尚未排序好的节点，则通过工作指针q从虚拟的头节点找到最后一个小于节点值，插入
// 需要注意的是不要断链
func insertionSortListCore(head *ListNode) *ListNode {
	p := head
	// 虚拟头节点
	head = &ListNode{
		Val:  math.MinInt32,
		Next: nil,
	}
	for p != nil {
		//为了不断链，将p从未排序好的结果中取下来
		q := p.Next
		p.Next = nil
		// 工作指针
		r := head
		// 找到最后一个小于其val的位置
		for r.Next != nil && r.Next.Val < p.Val {
			r = r.Next
		}
		// 在r后插入p
		p.Next = r.Next
		r.Next = p
		// 继续处理剩下未处理好的节点
		p = q
	}
	return head.Next
}

//另一条奇葩的思路是对值进行排序，然后填入链表中，但其实不是很符合题意。所以这里未给出实现
