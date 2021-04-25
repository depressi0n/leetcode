package question

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */
type Node133 struct {
	Val       int
	Neighbors []*Node133
}

func cloneGraph(node *Node133) *Node133 {
	//每个值都是独一无二的
	//做一个map
	m := make(map[int]*Node133)
	//用dfs进行复制
	var dfs func(*Node133) *Node133
	dfs = func(node *Node133) *Node133 {
		if _, ok := m[node.Val]; !ok { //之前没有复制过
			m[node.Val] = &Node133{
				Val:       node.Val,
				Neighbors: make([]*Node133, len(node.Neighbors)),
			}
		}
		for i := 0; i < len(node.Neighbors); i++ {
			m[node.Val].Neighbors[i] = dfs(node.Neighbors[i])
		}
		return node
	}
	return dfs(node)
}
