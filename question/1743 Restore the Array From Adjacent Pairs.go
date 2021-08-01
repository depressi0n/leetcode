package question

func Answer1743(adjacentPairs [][]int) []int {
	return restoreArray(adjacentPairs)
}
func restoreArray(adjacentPairs [][]int) []int {
	// n个不同元素，相邻元素给出，共n-1对
	// 恢复出整个数组
	// 思路：开始节点只会有一个相邻元素，结尾元素也只会有一个相邻元素
	// 用一个map记录指定元素的相邻元素，用一个map记录是否已经使用过
	res := make([]int, len(adjacentPairs)+1)
	m1 := make(map[int][]int, len(adjacentPairs)+1)
	m2 := make(map[int]bool, len(adjacentPairs)+1)
	for i := 0; i < len(adjacentPairs); i++ {
		if _, ok := m1[adjacentPairs[i][0]]; !ok {
			m1[adjacentPairs[i][0]] = make([]int, 0, 2)
			m1[adjacentPairs[i][0]] = append(m1[adjacentPairs[i][0]], adjacentPairs[i][1])
		} else {
			m1[adjacentPairs[i][0]] = append(m1[adjacentPairs[i][0]], adjacentPairs[i][1])
		}
		if _, ok := m1[adjacentPairs[i][1]]; !ok {
			m1[adjacentPairs[i][1]] = make([]int, 0, 2)
			m1[adjacentPairs[i][1]] = append(m1[adjacentPairs[i][1]], adjacentPairs[i][0])
		} else {
			m1[adjacentPairs[i][1]] = append(m1[adjacentPairs[i][1]], adjacentPairs[i][0])
		}
		if _, ok := m2[adjacentPairs[i][0]]; !ok {
			m2[adjacentPairs[i][0]] = true
		}
		if _, ok := m2[adjacentPairs[i][1]]; !ok {
			m2[adjacentPairs[i][1]] = true
		}
	}
	flag := false
	for item, nums := range m1 {
		if len(nums) == 1 {
			if !flag {
				res[0] = item
				m2[item] = false
				flag = true
			} else {
				res[len(adjacentPairs)] = item
				m2[item] = false
				break
			}
		}
	}
	index := 1
	res[index] = m1[res[0]][0]
	m2[res[index]] = false
	for index < len(adjacentPairs)-1 {
		// 已经使用过
		if !m2[m1[res[index]][0]] {
			res[index+1] = m1[res[index]][1]
			m2[m1[res[index]][1]] = false
		} else {
			res[index+1] = m1[res[index]][0]
			m2[m1[res[index]][0]] = false
		}
		index++
	}
	return res
}

func restoreArray1(adjacentPairs [][]int) []int {
	// n个不同元素，相邻元素给出，共n-1对
	// 恢复出整个数组
	// 思路：开始节点只会有一个相邻元素，结尾元素也只会有一个相邻元素
	// 用一个map记录指定元素的相邻元素，用一个map记录是否已经使用过
	n := len(adjacentPairs) + 1
	g := make(map[int][]int, n)
	for _, p := range adjacentPairs {
		v, w := p[0], p[1]
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	ans := make([]int, n)
	for e, adj := range g {
		if len(adj) == 1 {
			ans[0] = e
			break
		}
	}

	ans[1] = g[ans[0]][0]
	for i := 2; i < n; i++ {
		adj := g[ans[i-1]]
		if ans[i-2] == adj[0] {
			ans[i] = adj[1]
		} else {
			ans[i] = adj[0]
		}
	}
	return ans
}
