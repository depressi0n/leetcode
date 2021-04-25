package question

func IsBipartite(graph [][]int) bool {
	//假定没有自环，无向图，可能不是连通图
	//二分图定义：每条边的两个端点分别在一个集合中
	//主要思想：从任意一条边开始，进行集合并
	//第一行的非零值开始，进行dfs，将两个集合不断扩充
	// 知道所有的边都遍历完毕，不冲突则说明是二分图
	setA := make(map[int]struct{}, len(graph))
	setB := make(map[int]struct{}, len(graph))
	visit := make(map[int]struct{}, len(graph))
	var insert = func(set int, val int) int {
		switch set {
		// 放入之前检查
		case 0:
			// 如果被插入的值已经在B中了，报错
			if _, ok := setB[val]; ok {
				return -1
			}
			setA[val] = struct{}{}
		case 1:
			// 如果被插入的值已经在A中了，报错
			if _, ok := setA[val]; ok {
				return -1
			}
			setB[val] = struct{}{}
		}
		return 0
	}
	queue := make([]int, 0, len(graph))
	queue = append(queue, 0)
	for len(queue) != 0 || len(visit) != len(graph) {
		if len(queue) == 0 {
			// 从没有visit数组中找一个出来
			for i := 0; i < len(graph); i++ {
				if _, ok := visit[i]; !ok {
					queue = append(queue, i)
					break
				}
			}
			continue
		}
		i := queue[0]
		visit[i] = struct{}{}
		queue = queue[1:]
		var addedSet int
		// 当前的值已经在set B中，则与他相连的其他端点应该放入A中
		// 检查当前值在哪一个set中
		if _, ok := setA[i]; ok {
			//放入B中
			addedSet = 1
		} else if _, ok = setB[i]; ok {
			addedSet = 0
		} else {
			setA[i] = struct{}{}
			addedSet = 1
		}
		for j := 0; j < len(graph[i]); j++ {
			if insert(addedSet, graph[i][j]) != -1 {
				if _, ok := visit[graph[i][j]]; !ok {
					queue = append(queue, graph[i][j])
				}
			} else {
				return false
			}
		}
	}
	return true
}

// TODO: 优化，方向是直接BFS遍历，然后检查是否有冲突，最好是用并查集来做
func IsBipartite2(graph [][]int) bool {
	visit := make(map[int]int, len(graph)) //映射顶点在BFS序列中的位置
	queue := make([]int, len(graph))
	start, end := 0, 0
	mark := make([]int, len(graph))
	for {
		if start == end {
			for i := 0; i < len(graph); i++ {
				if _, ok := visit[i]; !ok { //没有记录过
					queue[end] = i
					visit[i] = end
					mark[end] = 0 // 放入左边数组中
					end++
					break
				}
			}
		}
		if start == end {
			return true
		}
		cur := queue[start]
		for i := 0; i < len(graph[cur]); i++ {
			// 查询另一个端点是否访问过，如果访问过肯定会记录标记
			// 检查标记是否冲突
			loc, ok := visit[graph[cur][i]]
			if ok {
				// 检查是否冲突
				if mark[loc]^mark[start] == 0 {
					return false
				}
			} else {
				queue[end] = graph[cur][i]
				visit[graph[cur][i]] = end //标记已经放入队列
				mark[end] = mark[start] ^ 1
				end++
			}
		}
		start++
	}
}
func IsBipartite3(graph [][]int) bool {
	queue := make([]int, len(graph))
	start, end := 0, 0
	mark := make([]int, len(graph))
	for i := 0; i < len(graph); i++ {
		mark[i] = 2
	}
	for {
		if start == end {
			for i := 0; i < len(graph); i++ {
				if mark[i] == 2 { //没有记录过
					queue[end] = i
					mark[i] = 0 // 放入左边数组中
					end++
					break
				}
			}
		}
		if start == end {
			return true
		}
		cur := queue[start]
		for i := 0; i < len(graph[cur]); i++ {
			// 查询另一个端点是否访问过，如果访问过肯定会记录标记
			// 检查标记是否冲突
			switch mark[graph[cur][i]] {
			case 2:
				queue[end] = graph[cur][i]
				end++
				mark[graph[cur][i]] = mark[cur] ^ 1
			case 1, 0:
				// 检查
				if mark[graph[cur][i]]^mark[cur] != 1 {
					return false
				}
			}
		}
		start++
	}
}
