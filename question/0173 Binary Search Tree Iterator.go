package question

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 像线索二叉树，用morris遍历肯定行
// 平均下来是O(1)
type BSTIterator struct {
	//1
	//cur int
	//inOrders   []int
	//2 模拟一个栈
	s   []*TreeNode
	cur *TreeNode
}

func BSTIteratorConstructor(root *TreeNode) BSTIterator {
	// 1.中序遍历
	//s:=make([]*TreeNode,0)
	//inOrder:=make([]int,0)
	//p:=root
	//for len(s)!=0 || p!=nil{
	//	if p!=nil{
	//		s=append(s,p)
	//		p=p.Left
	//	}else{
	//		p=s[len(s)-1]
	//		s=s[:len(s)-1]
	//		inOrder=append(inOrder,p.Val)
	//		p=p.Right
	//	}
	//}
	////for i:=0;i<len(inOrder);i++{
	////	fmt.Printf("%d->",inOrder[i])
	////}
	////fmt.Println()
	//return BSTIterator{
	//	cur: -1,
	//	inOrders: inOrder,
	//}
	// 2. 受控遍历
	s := make([]*TreeNode, 0)
	p := root
	return BSTIterator{
		s:   s,
		cur: p,
	}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	//1
	//this.cur++
	//return this.inOrders[this.cur]
	//2
	var res int
	flag := false
	for !flag && (len(this.s) != 0 || this.cur != nil) {
		if this.cur != nil {
			this.s = append(this.s, this.cur)
			this.cur = this.cur.Left
		} else {
			this.cur = this.s[len(this.s)-1]
			this.s = this.s[:len(this.s)-1]
			res = this.cur.Val
			flag = true
			this.cur = this.cur.Right
		}
	}
	return res
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	//1
	//if this.cur<len(this.inOrders)-1{
	//	return true
	//}
	//return false
	//2
	if len(this.s) == 0 && this.cur == nil {
		return false
	}
	return true
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

// TODO: morris 遍历的解法
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator1 struct {
	root *TreeNode
}

func Constructor1(root *TreeNode) BSTIterator1 {
	return BSTIterator1{root: root}
}

/** @return the next smallest number */
func (this *BSTIterator1) Next() int {
	i, _ := this.n(true)
	return i
}

/** @return whether we have a next smallest number */
func (this *BSTIterator1) HasNext() bool {
	_, ok := this.n(false)
	return ok
}

func (this *BSTIterator1) n(move bool) (int, bool) {
	var max *TreeNode

	for this.root != nil {
		if this.root.Left == nil {
			if move {
				defer func() {
					this.root = this.root.Right
				}()
			}
			return this.root.Val, true

		} else {
			//寻找左树最大节点
			max = this.root.Left
			for max.Right != nil {
				max = max.Right
			}

			//最大节点指向root
			max.Right = this.root

			//root = root.Left :移动root到root.left
			//root.Left = nil  :砍左子树，避免下一次遍历到root时，再进入到左子树
			this.root, this.root.Left = this.root.Left, nil
		}
	}
	return 0, false
}
