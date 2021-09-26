package question

// 给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。
// TODO:可以弄一个pool，减少内存分配
func generateTrees(n int) []*TreeNode {
	return generateTreesCore(n)
}

// 主要思想：首先明确当n==1时，就只有1种，当n==2时，有2种，当n==3时，有左1右1，左0右2，左2右0，共5种
// 此时，根据n的值，设定1-n的数组
// 使用不同的值作为root，递归建立子树
// 递归函数可设置为left，right，如果left>right，则直接返回nil，否则以其中一个作为root，继续
func generateTreesCore(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	nums := make([]int, n+1)
	for i := 1; i <= n; i++ {
		nums[i] = i
	}
	var generateSubTree func(int, int) []*TreeNode
	generateSubTree = func(left, right int) []*TreeNode {
		if left > right {
			return []*TreeNode{nil}
		}
		res := make([]*TreeNode, 0)
		for i := left; i <= right; i++ {
			leftTrees := generateSubTree(left, i-1)
			rightTrees := generateSubTree(i+1, right)
			for j := 0; j < len(leftTrees); j++ {
				for k := 0; k < len(rightTrees); k++ {
					root := &TreeNode{
						Val:   nums[i],
						Left:  nil,
						Right: nil,
					}
					root.Left = leftTrees[j]
					root.Right = rightTrees[k]
					res = append(res, root)
				}
			}
		}
		return res
	}
	//递归选择root
	res := generateSubTree(1, n)
	return res
}
