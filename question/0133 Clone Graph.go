package question

// 给你无向 连通 图中一个节点的引用，请你返回该图的 深拷贝（克隆）。
//图中的每个节点都包含它的值 val（int） 和其邻居的列表（list[Node]）。
// 提示：
//节点数不超过 100 。
//每个节点值Node.val 都是唯一的，1 <= Node.val <= 100。
//无向图是一个简单图，这意味着图中没有重复的边，也没有自环。
//由于图是无向的，如果节点 p 是节点 q 的邻居，那么节点 q 也必须是节点 p的邻居。
//图是连通图，你可以从给定节点访问到所有节点。

type Node133 struct {
	Val       int
	Neighbors []*Node133
}

func cloneGraph(node *Node133) *Node133 {
	return cloneGraphCore(node)
}

// 主要思想：根据bfs进行克隆，使用一个map记录已经克隆过的节点，如果遍历到就直接使用，否则进入下一层
func cloneGraphCore(node *Node133) *Node133 {
	if node == nil{
		return nil
	}
	//因为每个值都是独一无二的
	//使用一个map记录已经克隆过的节点
	m := make(map[int]*Node133)
	q := make([]*Node133, 0, 100)
	q = append(q, node)
	if _, ok := m[node.Val]; !ok {
		m[node.Val] = &Node133{
			Val:       node.Val,
			Neighbors: make([]*Node133, 0,len(node.Neighbors)),
		}
	}
	for len(q) > 0 {
		// 将没有入队的邻居入队
		for i := 0; i < len(q[0].Neighbors); i++ {
			// 没有进入过，则加入
			if _, ok := m[q[0].Neighbors[i].Val]; !ok {
				m[q[0].Neighbors[i].Val] = &Node133{
					Val:       q[0].Neighbors[i].Val,
					Neighbors: make([]*Node133, 0, len(q[0].Neighbors[i].Neighbors)),
				}
				q = append(q, q[0].Neighbors[i])
			}
			m[q[0].Val].Neighbors = append(m[q[0].Val].Neighbors, m[q[0].Neighbors[i].Val])
		}
		q = q[1:]
	}
	return m[node.Val]
}
